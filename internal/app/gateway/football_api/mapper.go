package football_api

import (
	"strconv"
	"time"

	match_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/match"
)

func mapMatches(
	matches []MatchDTO,
) ([]match_usecase.ExternalMatchOutput, error) {

	items := make([]match_usecase.ExternalMatchOutput, 0, len(matches))

	for _, match := range matches {
		startsAt, err := time.Parse(time.RFC3339, match.Fixture.Date)
		if err != nil {
			return nil, err
		}

		items = append(items, match_usecase.ExternalMatchOutput{
			ExternalID: strconv.FormatInt(match.Fixture.ID, 10),
			Stage:      normalizeStage(match.League.Round),
			HomeTeam: &match_usecase.ExternalTeamOutput{
				ExternalID: strconv.FormatInt(match.Teams.Home.ID, 10),
				Name:       match.Teams.Home.Name,
				Code:       match.Teams.Home.Code,
				FlagURL:    match.Teams.Home.Logo,
			},
			AwayTeam: &match_usecase.ExternalTeamOutput{
				ExternalID: strconv.FormatInt(match.Teams.Away.ID, 10),
				Name:       match.Teams.Away.Name,
				Code:       match.Teams.Away.Code,
				FlagURL:    match.Teams.Away.Logo,
			},
			StartsAt:  startsAt,
			HomeScore: match.Goals.Home,
			AwayScore: match.Goals.Away,
			Status:    normalizeStatus(match.Fixture.Status.Short),
		})
	}

	return items, nil
}
