package football_api

import (
	"strconv"
	"time"

	match_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/match"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/pkg/translations"
)

func mapMatches(
	matches []MatchDTO,
) ([]match_usecase.ExternalMatchOutput, error) {

	items := make([]match_usecase.ExternalMatchOutput, 0, len(matches))

	for _, match := range matches {
		startsAt, err := time.Parse(time.RFC3339, match.UTCDate)
		if err != nil {
			return nil, err
		}

		var homeTeam *match_usecase.ExternalTeamOutput
		if match.HomeTeam.ID != nil {
			homeTeam = &match_usecase.ExternalTeamOutput{
				ExternalID: strconv.FormatInt(*match.HomeTeam.ID, 10),
				Name: translations.TranslateTeamName(
					safeString(match.HomeTeam.Name),
				),
				Code:    safeString(match.HomeTeam.TLA),
				FlagURL: safeString(match.HomeTeam.Crest),
			}
		}

		var awayTeam *match_usecase.ExternalTeamOutput
		if match.AwayTeam.ID != nil {
			awayTeam = &match_usecase.ExternalTeamOutput{
				ExternalID: strconv.FormatInt(*match.AwayTeam.ID, 10),
				Name: translations.TranslateTeamName(
					safeString(match.AwayTeam.Name),
				),
				Code:    safeString(match.AwayTeam.TLA),
				FlagURL: safeString(match.AwayTeam.Crest),
			}
		}

		items = append(items, match_usecase.ExternalMatchOutput{
			ExternalID: strconv.FormatInt(match.ID, 10),

			Stage: normalizeStage(match.Stage),

			HomeTeam: homeTeam,
			AwayTeam: awayTeam,

			StartsAt: startsAt,

			HomeScore: match.Score.FullTime.Home,
			AwayScore: match.Score.FullTime.Away,

			Status: normalizeStatus(match.Status),
		})
	}

	return items, nil
}

func safeString(value *string) string {
	if value == nil {
		return ""
	}

	return *value
}
