package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/handler/schema"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest/response"
)

const userPattern = "/users"

func (h *Handler) userSetupRoutes(router chi.Router) {
	router.Route(userPattern, func(r chi.Router) {
		r.Get("/{id}", h.getUserByID())
	})
}

func (h *Handler) getUserByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &response.Response{}

		idParam := chi.URLParam(r, "id")

		id, err := uuid.Parse(idParam)
		if err != nil {
			resp = response.BadRequest(err, "invalid user id")
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		user, err := h.app.NewUserUseCase(h.app.DB.Pool).GetUserByUUID(r.Context(), id)
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		resp = response.OK(schema.UserResponseFromEntity(user))
		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
	}
}
