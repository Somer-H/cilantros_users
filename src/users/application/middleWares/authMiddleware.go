package middlewares

import (
	"net/http"
	"strings"
	repository "users_api/src/users/application/reposoitory"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	TokenManager repository.TokenManager
}

func NewAuthMiddleware(tokenManager repository.TokenManager) *AuthMiddleware {
	return &AuthMiddleware{TokenManager: tokenManager}
}

func (m *AuthMiddleware) Handle(c *gin.Context) {
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid Authorization header"})
        c.Abort()
        return
    }
    tokenString := strings.TrimPrefix(authHeader, "Bearer ")
    valid, claims, err := m.TokenManager.ValidateToken(tokenString)
    if err != nil || !valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
        c.Abort()
        return
    }
    username, ok := claims["username"].(string)
    if ok {
        c.Set("username", username)
    }
    c.Next()
}