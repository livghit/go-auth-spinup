package models

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt"
)

type JWTConfigs struct {
	SingMethod jwt.SigningMethod
	Claims     jwt.Claims
}

type JWT struct {
	Value string
}

var (
	key   []byte
	token *jwt.Token
	s     string
)

func CreateToken(configs JWTConfigs) (string, error) {
	key = []byte(os.Getenv("APP_KEY"))
	if key == nil {
		panic("NO ENCRYPTION KEY FOUND")
	}

	if configs.Claims != nil {
		token = jwt.NewWithClaims(configs.SingMethod, configs.Claims)
		log.Println("Token created with claims")
	} else {
		token = jwt.New(configs.SingMethod)
		log.Println("Token created without claims")
	}

	s, err := token.SignedString(key)
	if err != nil {
		return s, err
	}
	return s, nil
}

func ValidateJWT(jwt string) {
	// perform jwt token validation
}
