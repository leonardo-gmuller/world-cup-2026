package schema

import prediction_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/prediction"

type RankingItemResponse struct {
	UserID           int64   `json:"user_id"`
	UserUUID         string  `json:"user_uuid"`
	Name             string  `json:"name"`
	TotalPoints      float64 `json:"total_points"`
	PredictionsCount int64   `json:"predictions_count"`
	Position         int     `json:"position"`
}

func RankingListResponseFromUseCase(
	ranking []prediction_usecase.RankingItemOutput,
) []RankingItemResponse {
	items := make([]RankingItemResponse, 0, len(ranking))

	for index, item := range ranking {
		items = append(items, RankingItemResponse{
			UserID:           item.UserID,
			UserUUID:         item.UserUUID.String(),
			Name:             item.Name,
			TotalPoints:      item.TotalPoints,
			PredictionsCount: item.PredictionsCount,
			Position:         index + 1,
		})
	}

	return items
}
