package match_usecase

import (
	"context"
	"strconv"
	"time"
)

func (u *MatchUseCase) SyncLiveResults(ctx context.Context) error {
	matchesToSync, err := u.repo.ListMatchesToSyncLiveResults(ctx)
	if err != nil {
		return err
	}

	if len(matchesToSync) == 0 {
		return nil
	}

	dates := make(map[string]time.Time)

	for _, match := range matchesToSync {
		date := match.StartsAt.UTC()
		dateKey := date.Format("2006-01-02")

		if _, exists := dates[dateKey]; !exists {
			dates[dateKey] = date
		}
	}

	for _, date := range dates {
		externalMatches, err := u.liveClient.FetchTodayMatches(ctx, date)
		if err != nil {
			return err
		}

		for _, externalMatch := range externalMatches {
			if externalMatch.HomeTeam == nil || externalMatch.AwayTeam == nil {
				continue
			}

			internalMatch, err := u.repo.FindMatchForLiveSync(
				ctx,
				externalMatch.StartsAt,
				externalMatch.HomeTeam.Name,
				externalMatch.AwayTeam.Name,
			)
			if err != nil {
				return err
			}

			if internalMatch == nil {
				continue
			}

			apiFootballID, err := strconv.ParseInt(externalMatch.ExternalID, 10, 64)
			if err != nil {
				return err
			}

			scoreChanged :=
				!sameIntPointer(internalMatch.HomeScore, externalMatch.HomeScore) ||
					!sameIntPointer(internalMatch.AwayScore, externalMatch.AwayScore)

			statusChanged := internalMatch.Status != externalMatch.Status

			if !scoreChanged && !statusChanged {
				continue
			}

			_, err = u.repo.UpdateLiveResult(
				ctx,
				internalMatch.ID,
				apiFootballID,
				externalMatch.HomeScore,
				externalMatch.AwayScore,
				externalMatch.Status,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func sameIntPointer(a *int, b *int) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	return *a == *b
}
