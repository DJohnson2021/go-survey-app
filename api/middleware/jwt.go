package middleware

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(/* user data, if needed */) (string, error) {
	// Set up JWT claims (these can include user data if needed)
	claims := jwt.MapClaims{
		"iss": "Go-Survey-App",
		"exp": time.Now().Add(time.Hour * 72).Unix(),  // Set expiration time, for example, 72 hours from now
		// ... add other claims as needed
	}

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	signedToken, err := token.SignedString([]byte("your_secret_key"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}