package middlewares

import (
	"strings"
	"time"

	"github.com/PradnyaKuswara/sniffcrape/internal/models"
	"github.com/PradnyaKuswara/sniffcrape/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var jwtSecret = []byte(viper.GetString("JWT_SECRET"))

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			status, message := utils.MapErrorToStatusCode(jwt.ErrTokenMalformed)
			utils.RespondWithError(c, status, message)
			c.Abort()
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, _ := jwt.ParseWithClaims(tokenStr, &models.JwtAttributes{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		claims, ok := token.Claims.(*models.JwtAttributes)
		if !ok || claims.ExpiresAt == nil || claims.ExpiresAt.Time.Before(time.Now()) {
			status, message := utils.MapErrorToStatusCode(jwt.ErrTokenExpired)
			utils.RespondWithError(c, status, message)
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("userEmail", claims.UserEmail)
		c.Set("username", claims.Username)
		c.Set("firstName", claims.FirstName)
		c.Set("lastName", claims.LastName)
		c.Set("avatarURL", claims.AvatarURL)
		c.Next()
	}
}
