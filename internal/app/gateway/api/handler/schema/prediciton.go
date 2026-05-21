package schema

import (
	"fmt"

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
