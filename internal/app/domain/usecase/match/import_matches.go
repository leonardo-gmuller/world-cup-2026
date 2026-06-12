package match_usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/constants"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

func (u *MatchUseCase) ImportMatches(ctx context.Context) error {
	externalMatches, err := u.client.FetchWorldCupMatches(ctx)
	if err != nil {
		return err
	}

	for _, externalMatch := range externalMatches {
		existingMatch, err := u.repo.GetMatchByExternalID(ctx, externalMatch.ExternalID)
		if err != nil {
			return err
		}

		if existingMatch != nil {
			switch existingMatch.Status {
			case constants.MatchStatusLive, constants.MatchStatusFinished:
				continue
			}
		}

		var homeTeamID *int64
		var awayTeamID *int64

		if externalMatch.HomeTeam != nil {
			team, err := u.repo.UpsertTeam(ctx, entity.Team{
				UUID:       uuid.New(),
				ExternalID: stringPointer(externalMatch.HomeTeam.ExternalID),
				Name:       externalMatch.HomeTeam.Name,
				ShortName:  stringPointer(externalMatch.HomeTeam.ShortName),
				Code:       stringPointer(externalMatch.HomeTeam.Code),
				FlagURL:    stringPointer(externalMatch.HomeTeam.FlagURL),
			})
			if err != nil {
				return err
			}

			homeTeamID = &team.ID
		}

		if externalMatch.AwayTeam != nil {
			team, err := u.repo.UpsertTeam(ctx, entity.Team{
				UUID:       uuid.New(),
				ExternalID: stringPointer(externalMatch.AwayTeam.ExternalID),
				Name:       externalMatch.AwayTeam.Name,
				ShortName:  stringPointer(externalMatch.AwayTeam.ShortName),
				Code:       stringPointer(externalMatch.AwayTeam.Code),
				FlagURL:    stringPointer(externalMatch.AwayTeam.FlagURL),
			})
			if err != nil {
				return err
			}

			awayTeamID = &team.ID
		}

		match := entity.Match{
			UUID:         uuid.New(),
			ExternalID:   stringPointer(externalMatch.ExternalID),
			Stage:        externalMatch.Stage,
			GroupName:    externalMatch.GroupName,
			HomeTeamID:   homeTeamID,
			AwayTeamID:   awayTeamID,
			HomeTeamName: getTeamName(externalMatch.HomeTeam),
			AwayTeamName: getTeamName(externalMatch.AwayTeam),
			StartsAt:     externalMatch.StartsAt,
			HomeScore:    externalMatch.HomeScore,
			AwayScore:    externalMatch.AwayScore,
			Status:       externalMatch.Status,
		}

		_, err = u.repo.UpsertMatch(ctx, match)
		if err != nil {
			return err
		}
	}

	return nil
}

func stringPointer(value string) *string {
	if value == "" {
		return nil
	}

	return &value
}

func getTeamName(team *ExternalTeamOutput) *string {
	if team == nil || team.Name == "" {
		return nil
	}

	return &team.Name
}
