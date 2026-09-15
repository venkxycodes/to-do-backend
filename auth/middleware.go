package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const claimsKey = "auth_claims"

func RequireAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
			return
		}
		claims, err := VerifyToken(secret, strings.TrimPrefix(header, "Bearer "))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}
		c.Set(claimsKey, claims)
		c.Next()
	}
}

func ClaimsFromContext(c *gin.Context) (Claims, bool) {
	value, exists := c.Get(claimsKey)
	if !exists {
		return Claims{}, false
	}
	claims, ok := value.(Claims)
	return claims, ok
}
