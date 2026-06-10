package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/handler/schema"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest/response"
)

const rankingPattern = "/ranking"

func (h *Handler) rankingSetupRoutes(router chi.Router) {
	router.Route(rankingPattern, func(r chi.Router) {
		r.Get("/group/{group_id}", h.getGroupRanking())
	})
}

func (h *Handler) getGroupRanking() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &response.Response{}

		groupIDParam := chi.URLParam(r, "group_id")

		stageParam := r.URL.Query().Get("stage")

		var stage *string
		if stageParam != "" {
			stage = &stageParam
		}

		usecase := h.app.NewPredictionUseCase(h.app.DB.Pool)

		ranking, err := usecase.GetGroupRanking(r.Context(), groupIDParam, stage)
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		resp = response.OK(schema.RankingListResponseFromUseCase(ranking))
		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
	}
}
