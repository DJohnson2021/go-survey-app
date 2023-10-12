package middleware

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/DJohnson2021/go-survey-app/db"
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
	oauthState := generateStateOauthCookies(c)
	u := oauthConfig.AuthCodeURL(oauthState)
	return c.Redirect(u)
}

func oauthGoogleCallBack(c *fiber.Ctx) error {
	oauthState := c.Cookies("oauthstate")

	if c.Query("state") != oauthState {
		log.Println("invalid oauth google state")
		return c.Redirect("/")
	}

	data, err := getUserDataFromGoogle(c.Query("code"))
	if err != nil {
		log.Println(err.Error())
		return c.Redirect("/")
	}
	
	// GetOrCreate User in your db.
	// Redirect or response with a token.
	// More code .....

	return c.SendString(fmt.Sprintf("UserInfo: %s\n", data))
}

func generateStateOauthCookies(c *fiber.Ctx) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	fiberCookie := &fiber.Cookie{
		Name: "oauthstate",
		Value: state,
		Expires: expiration,
	}

	c.Cookie(fiberCookie)

	return state
}


func getUserDataFromGoogle(code string) ([]byte, error) {
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("could not exchange authorization code into token: %s", err.Error())
	}

	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get user info %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %s", err.Error())
	}

	return contents, nil
}










