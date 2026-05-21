package cronjob

import (
	"context"
	"fmt"
)

func (h *Handler) ImportMatches(ctx context.Context) error {
	const operation = "Cronjob.Handler.ImportMatches"

	err := h.useCase.ImportMatches(ctx)
	if err != nil {
		return fmt.Errorf("%s -> %w", operation, err)
	}
	return nil
}
