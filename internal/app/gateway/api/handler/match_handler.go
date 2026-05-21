package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/handler/schema"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest/response"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres"
)

const matchPattern = "/matches"

func (h *Handler) matchSetupRoutes(router chi.Router) {
	router.Route(matchPattern, func(r chi.Router) {
		r.Get("/", h.listMatches())
		r.Get("/stage/{stage}", h.listMatchesByStage())
		r.Post("/import", h.importMatches())
	})
}

func (h *Handler) listMatches() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &response.Response{}

		usecase := h.app.NewMatchUseCase(h.app.DB.Pool)

		matches, err := usecase.ListMatches(r.Context())
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		resp = response.OK(schema.MatchListResponseFromEntity(matches))
		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
	}
}

func (h *Handler) listMatchesByStage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &response.Response{}

		stage := chi.URLParam(r, "stage")

		usecase := h.app.NewMatchUseCase(h.app.DB.Pool)

		matches, err := usecase.ListMatchesByStage(r.Context(), stage)
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		resp = response.OK(schema.MatchListResponseFromEntity(matches))
		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
	}
}

func (h *Handler) importMatches() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &response.Response{}

		err := h.app.DB.WithTx(r.Context(), func(ctx context.Context, dbtx postgres.DBTX) error {
			usecase := h.app.NewMatchUseCase(dbtx)
			return usecase.ImportMatches(ctx)
		})

		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		resp = response.OK(map[string]string{
			"message": "matches imported successfully",
		})
		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
	}
}
