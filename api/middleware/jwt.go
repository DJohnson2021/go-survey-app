package middleware

import (
	"log"
	"time"
	"github.com/DJohnson2021/go-survey-app/utils"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"isAdmin"`
	jwt.StandardClaims
}

func GenerateJWT(name, email string) (string, error) {
	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &Claims{
		Name:    name,
		Email:   email,
		IsAdmin: false,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	jwtKey, err := utils.GetJWTSecret()
	if err != nil {
		log.Fatalf("Error getting jwt secret: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed_token, err := token.SignedString(jwtKey)
	if err != nil {
		log.Fatalf("error signing token: %v", err)
	}

	return signed_token, nil
}
