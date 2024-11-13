package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

func CreateToken(userID int64, userName string, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   userID,
		"user_name": userName,
		"exp":       time.Now().Add(time.Minute * 1).Unix(),
		"iat":       time.Now().Unix(),
	})

	key := []byte(secretKey)

	return token.SignedString(key)
}

func ValidateToken(tokenString string, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to parse token")
		return 0, "", errors.New("failed to parse token")
	}

	if !token.Valid {
		log.Error().Msg("invalid token")
		return 0, "", errors.New("invalid token")
	}

	return int64(claims["user_id"].(float64)), claims["user_name"].(string), nil
}

func ValidateTokenWithoutExpired(tokenString string, secretKey string) (int64, string, error) {
	key := []byte(secretKey)
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}, jwt.WithoutClaimsValidation())
	if err != nil {
		log.Error().Err(err).Msg("failed to parse token")
		return 0, "", errors.New("failed to parse token")
	}

	if !token.Valid {
		log.Error().Msg("invalid token")
		return 0, "", errors.New("invalid token")
	}

	return int64(claims["user_id"].(float64)), claims["user_name"].(string), nil
}
