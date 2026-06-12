package api_football

import (
	"strconv"
	"time"

	match_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/match"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/pkg/translations"
)

func mapFixtures(fixtures []FixtureDTO) ([]match_usecase.ExternalMatchOutput, error) {
	output := make([]match_usecase.ExternalMatchOutput, 0, len(fixtures))

	for _, fixture := range fixtures {
		startsAt, err := time.Parse(time.RFC3339, fixture.Fixture.Date)
		if err != nil {
			return nil, err
		}

		output = append(output, match_usecase.ExternalMatchOutput{
			ExternalID: strconv.FormatInt(fixture.Fixture.ID, 10),
			Stage:      normalizeStage(fixture.League.Round),
			HomeTeam: &match_usecase.ExternalTeamOutput{
				ExternalID: strconv.FormatInt(fixture.Teams.Home.ID, 10),
				Name:       translations.TranslateTeamName(fixture.Teams.Home.Name),
				FlagURL:    fixture.Teams.Home.Logo,
			},
			AwayTeam: &match_usecase.ExternalTeamOutput{
				ExternalID: strconv.FormatInt(fixture.Teams.Away.ID, 10),
				Name:       translations.TranslateTeamName(fixture.Teams.Away.Name),
				FlagURL:    fixture.Teams.Away.Logo,
			},
			StartsAt:  startsAt,
			HomeScore: fixture.Goals.Home,
			AwayScore: fixture.Goals.Away,
			Status:    normalizeStatus(fixture.Fixture.Status.Short),
		})
	}

	return output, nil
}
