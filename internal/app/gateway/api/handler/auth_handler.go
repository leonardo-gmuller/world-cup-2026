package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/erring"
	auth_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/auth"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/handler/schema"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest/response"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres"
)

const authPattern = "/auth"

func (h *Handler) authSetupRoutes(router chi.Router) {
	router.Route(authPattern, func(r chi.Router) {
		r.Post("/register", h.register())
		r.Post("/login", h.login())
	})
}

func (h *Handler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &response.Response{}

		var req schema.RegisterRequest
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
			usecase := h.app.NewAuthUseCase(dbtx)

			out, err := usecase.Register(r.Context(), auth_usecase.RegisterInput{
				Name:     req.Name,
				Email:    req.Email,
				Password: req.Password,
			})
			if err != nil {
				return err
			}

			resp = response.Created(schema.AuthResponseFromUseCase(out))
			return nil
		})

		if err != nil {
			if errors.Is(err, erring.ErrEmailAlreadyUsed) {
				resp = response.BadRequest(err, err.Error())
				rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
				return
			}

			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)

	}
}

func (h *Handler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &response.Response{}

		var req schema.LoginRequest
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
			usecase := h.app.NewAuthUseCase(dbtx)

			out, err := usecase.Login(r.Context(), auth_usecase.LoginInput{
				Email:    req.Email,
				Password: req.Password,
			})
			if err != nil {
				return err
			}

			resp = response.OK(schema.AuthResponseFromUseCase(out))
			return nil
		})

		if err != nil {
			if errors.Is(err, erring.ErrInvalidCredentials) {
				resp = response.Unauthorized()
				rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
				return
			}

			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
	}
}
