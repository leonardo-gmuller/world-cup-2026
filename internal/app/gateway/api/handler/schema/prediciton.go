package schema

import (
	"fmt"
	"time"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/dto"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type SavePredictionRequest struct {
	GroupID   string `json:"group_id"`
	MatchID   string `json:"match_id"`
	HomeScore int    `json:"home_score"`
	AwayScore int    `json:"away_score"`
}

func (r *SavePredictionRequest) Validate() error {
	if r.GroupID == "" {
		return fmt.Errorf("group_id is required")
	}

	if r.MatchID == "" {
		return fmt.Errorf("match_id is required")
	}

	if r.HomeScore < 0 || r.AwayScore < 0 {
		return fmt.Errorf("scores cannot be negative")
	}

	return nil
}

type PredictionResponse struct {
	ID           string  `json:"id"`
	GroupID      string  `json:"group_id"`
	MatchID      string  `json:"match_id"`
	HomeScore    int     `json:"home_score"`
	AwayScore    int     `json:"away_score"`
	Points       float64 `json:"points"`
	Calculated   bool    `json:"calculated"`
	CalculatedAt *string `json:"calculated_at"`
	CreatedAt    string  `json:"created_at"`
}

func PredictionResponseFromEntity(prediction *entity.Prediction) PredictionResponse {
	var calculatedAt *string
	if prediction.CalculatedAt != nil {
		value := prediction.CalculatedAt.Format("2006-01-02T15:04:05Z07:00")
		calculatedAt = &value
	}

	return PredictionResponse{
		ID:           prediction.UUID.String(),
		GroupID:      prediction.GroupUUID.String(),
		MatchID:      prediction.MatchUUID.String(),
		HomeScore:    prediction.HomeScore,
		AwayScore:    prediction.AwayScore,
		Points:       prediction.Points,
		Calculated:   prediction.Calculated,
		CalculatedAt: calculatedAt,
		CreatedAt:    prediction.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func PredictionListResponseFromEntity(predictions []entity.Prediction) []PredictionResponse {
	items := make([]PredictionResponse, 0, len(predictions))

	for _, prediction := range predictions {
		items = append(items, PredictionResponseFromEntity(&prediction))
	}

	return items
}

type PredictionReminderResponse struct {
	MatchID         int64     `json:"match_id"`
	GroupID         int64     `json:"group_id"`
	GroupName       string    `json:"group_name"`
	HomeTeamName    string    `json:"home_team_name"`
	AwayTeamName    string    `json:"away_team_name"`
	HomeTeamFlagURL string    `json:"home_team_flag_url"`
	AwayTeamFlagURL string    `json:"away_team_flag_url"`
	StartsAt        time.Time `json:"starts_at"`
}

func PredictionReminderResponseFromDto(
	output dto.PredictionReminderOutput,
) PredictionReminderResponse {
	return PredictionReminderResponse{
		MatchID:         output.MatchID,
		GroupID:         output.GroupID,
		GroupName:       output.GroupName,
		HomeTeamName:    output.HomeTeamName,
		AwayTeamName:    output.AwayTeamName,
		HomeTeamFlagURL: output.HomeTeamFlagURL,
		AwayTeamFlagURL: output.AwayTeamFlagURL,
		StartsAt:        output.StartsAt,
	}
}

func PredictionReminderListResponseFromDto(
	outputs []dto.PredictionReminderOutput,
) []PredictionReminderResponse {
	items := make([]PredictionReminderResponse, 0, len(outputs))

	for _, output := range outputs {
		items = append(items, PredictionReminderResponseFromDto(output))
	}

	return items
}
