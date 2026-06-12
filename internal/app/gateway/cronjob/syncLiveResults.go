package cronjob

import (
	"context"
	"fmt"
)

func (h *Handler) SyncLiveResults(ctx context.Context) error {
	const operation = "Cronjob.Handler.SyncLiveResults"

	err := h.useCase.SyncLiveResults(ctx)
	if err != nil {
		return fmt.Errorf("%s -> %w", operation, err)
	}
	return nil
}
