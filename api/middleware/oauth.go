package middleware

import (
	"context"
	"log"
	"os"
	"encoding/base64"
	"crypto/rand"
	"fmt"
	"io"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


func LoadEnv() {
	// Load .env file
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}


var oauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8000/api/survey", // Change this to your application's callback URL
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	// Use the actual links for the string slice instead of only the keywords
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.name"},
	Endpoint:     google.Endpoint,
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func oauthGoogleLogin(c *fiber.Ctx) error {
	LoadEnv()
	oauthState := generateStateOauthCookies(c)
	u := oauthConfig.AuthCodeURL(oauthState)
	return c.Redirect(u)
}










