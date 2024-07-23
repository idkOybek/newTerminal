package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/idkOybek/internal/logger"
	"github.com/idkOybek/internal/models"
	"github.com/idkOybek/internal/utils"
)

type contextKey string

const userContextKey contextKey = "user"

// AuthMiddleware проверяет наличие и валидность JWT токена
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Authorization header required")
			logger.ErrorLogger.Println("Authorization header missing")
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Authorization header format must be Bearer {token}")
			logger.ErrorLogger.Println("Invalid Authorization header format")
			return
		}

		token := parts[1]
		claims, err := utils.ParseJWT(token)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
			logger.ErrorLogger.Printf("Invalid token: %v", err)
			return
		}

		user := &models.User{
			ID:       claims.ID,
			Username: claims.Username,
			IsAdmin:  claims.IsAdmin,
		}

		ctx := context.WithValue(r.Context(), userContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserFromContext извлекает пользователя из контекста запроса
func GetUserFromContext(ctx context.Context) (*models.User, bool) {
	user, ok := ctx.Value(userContextKey).(*models.User)
	return user, ok
}
