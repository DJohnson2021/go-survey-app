package middleware

import (
	"fmt"
	"github.com/DJohnson2021/go-survey-app/utils"
	"github.com/golang-jwt/jwt"
)

func userAuthorized(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		jwtKey, err := utils.GetJWTSecret()
		if err != nil {
			return nil, fmt.Errorf("error getting JWT secret: %v", err)
		}
		return jwtKey, nil
	})

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if err != nil {
		return nil, err
	}

	return claims, nil
}

func adminAuthorized(tokenString string) (*Claims, error) {
	claims, err := userAuthorized(tokenString)
	if err != nil {
		return nil, fmt.Errorf("error authorizing user: %v", err)
	}

	// Load admin names and emails from .env
	adminNames, adminEmails, err := utils.GetAdminNamesAndEmails()
	if err != nil {
		return nil, fmt.Errorf("error getting admin names and emails: %v", err)
	}

	// Check if the user is an admin
	claims.IsAdmin = false
	for i, name := range adminNames {
		if name == claims.Name && adminEmails[i] == claims.Email {
			claims.IsAdmin = true
			break
		}
	}

	if !claims.IsAdmin {
		return nil, fmt.Errorf("user is not an admin")
	}

	return claims, nil
}
