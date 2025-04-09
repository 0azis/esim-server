package pkg

import (
	"errors"
	"esim/config"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt interface {
	SignJwt(userID int) (string, error)
	ParseJwt(jwtToken string) (int, error)
}

type jwtBuilder struct {
	cfg config.Config
}

func NewJwtBuilder(cfg config.Config) Jwt {
	return jwtBuilder{cfg}
}

func (j jwtBuilder) SignJwt(userID int) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userID,
	}).SignedString(j.cfg.SecretKey)

	return token, err
}

func (j jwtBuilder) ParseJwt(jwtToken string) (int, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (any, error) {
		return j.cfg.SecretKey, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["id"].(int), nil
	} else {
		return 0, errors.New("no claims")
	}
}
