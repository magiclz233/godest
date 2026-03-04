package middleware

import (
	"strings"

	"godest/internal/transport/http/response"
	"godest/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtUtil *utils.JWTUtil) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			_ = c.Error(response.Unauthorized("authorization header is required"))
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			_ = c.Error(response.Unauthorized("authorization format must be Bearer {token}"))
			c.Abort()
			return
		}

		claims, err := jwtUtil.ParseToken(parts[1])
		if err != nil {
			_ = c.Error(response.Unauthorized("invalid or expired token"))
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Next()
	}
}
