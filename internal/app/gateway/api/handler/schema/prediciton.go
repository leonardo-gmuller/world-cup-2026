package schema

import (
	"fmt"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
)

type SavePredictionRequest struct {
	GroupID   int64 `json:"group_id"`
	MatchID   int64 `json:"match_id"`
	HomeScore int   `json:"home_score"`
	AwayScore int   `json:"away_score"`
}

func (r *SavePredictionRequest) Validate() error {
	if r.GroupID == 0 {
		return fmt.Errorf("group_id is required")
	}

	if r.MatchID == 0 {
		return fmt.Errorf("match_id is required")
	}

	if r.HomeScore < 0 || r.AwayScore < 0 {
		return fmt.Errorf("scores cannot be negative")
	}

	return nil
}

type PredictionResponse struct {
	ID           string  `json:"id"`
	GroupID      int64   `json:"group_id"`
	UserID       int64   `json:"user_id"`
	MatchID      int64   `json:"match_id"`
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
		GroupID:      prediction.GroupID,
		UserID:       prediction.UserID,
		MatchID:      prediction.MatchID,
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
