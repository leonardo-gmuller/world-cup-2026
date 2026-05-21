package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	group_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/group"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/handler/schema"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/middleware"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest/response"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres"
)

const groupPattern = "/groups"

func (h *Handler) groupSetupRoutes(router chi.Router) {
	router.Route(groupPattern, func(r chi.Router) {
		r.Post("/", h.createGroup())
		r.Get("/", h.listGroupsByUser())
		r.Post("/join", h.joinGroup())
	})
}

func (h *Handler) createGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &response.Response{}

		var req schema.CreateGroupRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			resp = response.BadRequest(err, fmt.Errorf("invalid payload: %w", err).Error())
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		if err := req.Validate(); err != nil {
			resp = response.BadRequest(err, err.Error())
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		authUser, ok := middleware.GetAuthUser(r.Context())
		if !ok {
			resp := response.Unauthorized()
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		err := h.app.DB.WithTx(r.Context(), func(ctx context.Context, dbtx postgres.DBTX) error {
			usecase := h.app.NewGroupUseCase(dbtx)

			out, err := usecase.CreateGroup(ctx, group_usecase.CreateGroupInput{
				Name:        req.Name,
				Description: req.Description,
				OwnerID:     authUser.ID,
			})
			if err != nil {
				return err
			}

			resp = response.Created(schema.GroupResponseFromEntity(out))
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return nil
		})

		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
	}
}

func (h *Handler) joinGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &response.Response{}

		var req schema.JoinGroupRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			resp = response.BadRequest(err, fmt.Errorf("invalid payload: %w", err).Error())
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		if err := req.Validate(); err != nil {
			resp = response.BadRequest(err, err.Error())
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		err := h.app.DB.WithTx(r.Context(), func(ctx context.Context, dbtx postgres.DBTX) error {
			usecase := h.app.NewGroupUseCase(dbtx)

			out, err := usecase.JoinGroup(ctx, group_usecase.JoinGroupInput{
				UserID:     req.UserID,
				InviteCode: req.InviteCode,
			})
			if err != nil {
				return err
			}

			resp = response.Created(schema.GroupMemberResponseFromEntity(out))
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return nil
		})

		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
	}
}

func (h *Handler) listGroupsByUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &response.Response{}

		authUser, ok := middleware.GetAuthUser(r.Context())
		if !ok {
			resp := response.Unauthorized()
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		usecase := h.app.NewGroupUseCase(h.app.DB.Pool)

		groups, err := usecase.ListGroupsByUserID(r.Context(), authUser.ID)
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		resp = response.OK(schema.GroupListResponseFromEntity(groups))
		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
	}
}
