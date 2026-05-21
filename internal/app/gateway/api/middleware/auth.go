package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest/response"
)

type contextKey string

const userContextKey contextKey = "auth_user"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil || claims == nil {
			resp := response.Unauthorized()
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		user, err := userFromClaims(claims)
		if err != nil {
			resp := response.Unauthorized()
			rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func userFromClaims(claims map[string]interface{}) (*entity.User, error) {
	var userID int64

	switch v := claims["user_id"].(type) {
	case float64:
		userID = int64(v)

	case int64:
		userID = v

	case int:
		userID = int64(v)

	default:
		return nil, fmt.Errorf("invalid user_id")
	}

	userUUID, ok := claims["uuid"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid uuid")
	}

	name, _ := claims["name"].(string)
	email, _ := claims["email"].(string)

	return &entity.User{
		ID:    userID,
		UUID:  uuid.MustParse(userUUID),
		Name:  name,
		Email: email,
	}, nil
}

func GetAuthUser(ctx context.Context) (*entity.User, bool) {
	user, ok := ctx.Value(userContextKey).(*entity.User)
	return user, ok
}
