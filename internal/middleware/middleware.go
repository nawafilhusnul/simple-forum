package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nawafilhusnul/forum/internal/configs"
	"github.com/nawafilhusnul/forum/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		header = strings.TrimSpace(header)

		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		token := strings.Split(header, " ")[1]

		userID, userName, err := jwt.ValidateToken(token, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("user_id", userID)
		c.Set("user_name", userName)

		c.Next()
	}
}

func AuthRefreshMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretJWT
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		header = strings.TrimSpace(header)

		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		token := strings.Split(header, " ")[1]

		userID, userName, err := jwt.ValidateTokenWithoutExpired(token, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("user_id", userID)
		c.Set("user_name", userName)
		c.Next()
	}
}
