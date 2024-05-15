package middleware

import (
	"net/http"

	"blog.com/pkg/lib"
	"github.com/gin-gonic/gin"
)

// AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Token")

		if authHeader == "" {
			response := lib.Error("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		_, err := lib.ValidateToken(authHeader)
		if err != nil {
			response := lib.Error("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

	}
}
