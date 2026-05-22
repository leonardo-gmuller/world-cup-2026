package cronjob

import (
	"context"
	"fmt"
)

func (h *Handler) CalculateMatchPredictions(ctx context.Context) error {
	const operation = "Cronjob.Handler.CalculateMatchPredictions"

	err := h.useCase.CalculateMatchPredictions(ctx)
	if err != nil {
		return fmt.Errorf("%s -> %w", operation, err)
	}
	return nil
}
