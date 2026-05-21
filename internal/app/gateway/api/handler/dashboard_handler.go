package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/middleware"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest/response"
)

const dashboardPattern = "/dashboard"

func (h *Handler) dashboardSetupRoutes(router chi.Router) {
	router.Route(dashboardPattern, func(r chi.Router) {
		r.Get("/", h.getHomeDashboard())
	})
}

func (h *Handler) getHomeDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &response.Response{}

		authUser, ok := middleware.GetAuthUser(r.Context())
		if !ok {
			resp := response.Unauthorized()
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		dashboard, err := h.app.NewDashboardUseCase(h.app.DB.Pool).GetDashboardData(r.Context(), authUser.ID)
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		resp = response.OK(dashboard)
		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)

	}
}
