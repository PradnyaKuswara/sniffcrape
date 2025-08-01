package utils

import (
	"time"

	"github.com/PradnyaKuswara/sniffcrape/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

var jwtKey = []byte(viper.GetString("JWT_SECRET"))

func GenerateJWT(payload models.JwtAttributes) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         payload.ID,
		"email":      payload.Email,
		"username":   payload.Username,
		"first_name": payload.FirstName,
		"last_name":  payload.LastName,
		"avatar_url": payload.AvatarURL,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	})
	return token.SignedString(jwtKey)
}

func GetJWTSecret() []byte {
	return jwtKey
}
