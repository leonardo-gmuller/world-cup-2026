package schema

import (
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type MatchResponse struct {
	ID              string  `json:"id"`
	ExternalID      *string `json:"external_id"`
	Stage           string  `json:"stage"`
	GroupName       *string `json:"group_name"`
	HomeTeamID      *int64  `json:"home_team_id"`
	AwayTeamID      *int64  `json:"away_team_id"`
	HomeTeamName    *string `json:"home_team_name"`
	AwayTeamName    *string `json:"away_team_name"`
	StartsAt        string  `json:"starts_at"`
	HomeScore       *int    `json:"home_score"`
	AwayScore       *int    `json:"away_score"`
	Status          string  `json:"status"`
	HomeTeamFlagURL *string `json:"home_team_flag_url"`
	AwayTeamFlagURL *string `json:"away_team_flag_url"`
}

func MatchResponseFromEntity(match *entity.Match) MatchResponse {
	return MatchResponse{
		ID:              match.UUID.String(),
		ExternalID:      match.ExternalID,
		Stage:           match.Stage,
		GroupName:       match.GroupName,
		HomeTeamID:      match.HomeTeamID,
		AwayTeamID:      match.AwayTeamID,
		HomeTeamName:    match.HomeTeamName,
		AwayTeamName:    match.AwayTeamName,
		StartsAt:        match.StartsAt.Format("2006-01-02T15:04:05Z07:00"),
		HomeScore:       match.HomeScore,
		AwayScore:       match.AwayScore,
		Status:          match.Status,
		HomeTeamFlagURL: match.HomeTeamFlagURL,
		AwayTeamFlagURL: match.AwayTeamFlagURL,
	}
}

func MatchListResponseFromEntity(matches []entity.Match) []MatchResponse {
	items := make([]MatchResponse, 0, len(matches))

	for _, match := range matches {
		items = append(items, MatchResponseFromEntity(&match))
	}

	return items
}
