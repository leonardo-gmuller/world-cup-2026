package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	prediction_usecase "github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/usecase/prediction"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/handler/schema"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/middleware"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest/response"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/postgres"
)

const predictionPattern = "/predictions"

func (h *Handler) predictionSetupRoutes(router chi.Router) {
	router.Route(predictionPattern, func(r chi.Router) {
		r.Post("/", h.savePrediction())
		r.Get("/group/{group_id}", h.listUserPredictionsByGroup())
		r.Post("/match/{match_id}/calculate", h.calculateMatchPredictions())
	})
}

func (h *Handler) savePrediction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &response.Response{}

		authUser, ok := middleware.GetAuthUser(r.Context())
		if !ok {
			resp = response.Unauthorized()
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		var req schema.SavePredictionRequest
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
			usecase := h.app.NewPredictionUseCase(dbtx)

			out, err := usecase.SavePrediction(ctx, prediction_usecase.SavePredictionInput{
				GroupID:   req.GroupID,
				UserID:    authUser.ID,
				MatchID:   req.MatchID,
				HomeScore: req.HomeScore,
				AwayScore: req.AwayScore,
			})
			if err != nil {
				return err
			}

			response.Created(schema.PredictionResponseFromEntity(out))
			return nil
		})

		if err != nil {
			resp = response.BadRequest(err, err.Error())
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
	}
}

func (h *Handler) listUserPredictionsByGroup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &response.Response{}

		authUser, ok := middleware.GetAuthUser(r.Context())
		if !ok {
			resp = response.Unauthorized()
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		groupIDParam := chi.URLParam(r, "group_id")

		groupID, err := strconv.ParseInt(groupIDParam, 10, 64)
		if err != nil {
			resp = response.BadRequest(err, "invalid group_id")
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		usecase := h.app.NewPredictionUseCase(h.app.DB.Pool)

		predictions, err := usecase.ListPredictionsByUserAndGroup(
			r.Context(),
			groupID,
			authUser.ID,
		)
		if err != nil {
			resp = response.InternalServerError(err)
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		resp = response.OK(schema.PredictionListResponseFromEntity(predictions))
		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
	}
}

func (h *Handler) calculateMatchPredictions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := &response.Response{}

		matchIDParam := chi.URLParam(r, "match_id")

		matchID, err := strconv.ParseInt(matchIDParam, 10, 64)
		if err != nil {
			resp = response.BadRequest(err, "invalid match_id")
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		err = h.app.DB.WithTx(r.Context(), func(ctx context.Context, dbtx postgres.DBTX) error {
			usecase := h.app.NewPredictionUseCase(dbtx)
			return usecase.CalculateMatchPredictions(ctx, matchID)
		})

		if err != nil {
			resp = response.BadRequest(err, err.Error())
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		resp = response.OK(map[string]string{
			"message": "predictions calculated successfully",
		})
		rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
	}
}
