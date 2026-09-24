package middleware

import (
	"net/http"
	"strings"

	"api-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	CtxUserIDKey   = "userID"
	CtxUsernameKey = "username"
)

func AuthMiddleware(authUsecase usecase.AuthUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			SendError(c, http.StatusUnauthorized, "Authorization header is required")
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			SendError(c, http.StatusUnauthorized, "Invalid authorization header format (must be Bearer token)")
			return
		}

		tokenStr := parts[1]
		claims, err := authUsecase.ValidateToken(tokenStr)
		if err != nil {
			SendError(c, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		userUUID, err := uuid.Parse(claims.UserID)
		if err != nil {
			SendError(c, http.StatusUnauthorized, "Invalid user ID in token")
			return
		}

		c.Set(CtxUserIDKey, userUUID)
		c.Set(CtxUsernameKey, claims.Username)
		c.Next()
	}
}

func GetCurrentUserID(c *gin.Context) uuid.UUID {
	val, ok := c.Get(CtxUserIDKey)
	if !ok {
		return uuid.Nil
	}
	return val.(uuid.UUID)
}
