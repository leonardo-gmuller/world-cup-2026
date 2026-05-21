package handler

import (
	"net/http"
	"strconv"

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

		groupID, err := strconv.ParseInt(groupIDParam, 10, 64)
		if err != nil {
			resp = response.BadRequest(err, "invalid group_id")
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		usecase := h.app.NewPredictionUseCase(h.app.DB.Pool)

		ranking, err := usecase.GetGroupRanking(r.Context(), groupID)
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		resp = response.OK(schema.RankingListResponseFromUseCase(ranking))
		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
	}
}
