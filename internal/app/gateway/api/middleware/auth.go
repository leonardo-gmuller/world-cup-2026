package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/leonardo-gmuller/world-cup-2026/internal/app/domain/entity"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest"
	"github.com/leonardo-gmuller/world-cup-2026/internal/app/gateway/api/rest/response"
)

type contextKey string

const userContextKey contextKey = "auth_user"

type jwtService interface {
	Validate(token string) (*entity.User, error)
}

func Auth(jwt jwtService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				resp := response.Unauthorized()
				rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			if token == authHeader || token == "" {
				resp := response.Unauthorized()
				rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
				return
			}

			user, err := jwt.Validate(token)
			if err != nil {
				resp := response.Unauthorized()
				rest.SendJSON(w, resp.Status, resp.Payload, resp.Headers)
				return
			}

			ctx := context.WithValue(r.Context(), userContextKey, user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetAuthUser(ctx context.Context) (*entity.User, bool) {
	user, ok := ctx.Value(userContextKey).(*entity.User)
	return user, ok
}
