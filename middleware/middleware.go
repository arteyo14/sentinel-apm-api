package middleware

import (
	"net/http"
	"sentinel-apm-api/utils/context"
	"sentinel-apm-api/utils/response"
	"sentinel-apm-api/utils/token"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			response.HandleResponse(c, response.Response{
				Status: false,
				Code:   http.StatusUnauthorized,
				Error:  "Unauthorized",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.HandleResponse(c, response.Response{
				Status: false,
				Code:   http.StatusUnauthorized,
				Error:  "Unauthorized",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		claims, err := token.NewTokenManager().ValidateToken(tokenString)
		if err != nil {
			response.HandleResponse(c, response.Response{
				Status: false,
				Code:   http.StatusUnauthorized,
				Error:  "Invalid token",
			})
			c.Abort()
			return
		}

		context.SetSessionUser(context.SessionUser{
			UserID:    claims.UserID,
			Email:     claims.Email,
			FirstName: claims.FirstName,
			LastName:  claims.LastName,
		})
		c.Next()
	}
}
