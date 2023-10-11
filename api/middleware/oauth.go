package middleware

import (
	"context"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func init() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

var clientID = os.Getenv("GOOGLE_CLIENT_ID")
var clientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")

var (
	oauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/api/survey", // Change this to your application's callback URL
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint:     google.Endpoint,
	}
)

func HandleGoogleAuth(c *fiber.Ctx) error {
	// Generate URL and redirect to Google's OAuth page
	url := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

func HandleGoogleCallback(c *fiber.Ctx) error {
	// Get the code from the callback request
	code := c.Query("code")

	// Use the code to get the token from Google
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(400).SendString("Failed to exchange token: " + err.Error())
	}

	client := oauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		// Handle error
	}
	defer resp.Body.Close()

	// Decode the response to get user details
	var userInfo map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		// Handle error
	}

	// Extract relevant data
	email := userInfo["email"].(string)
	name := userInfo["name"].(string)
	password := userInfo["password"].(string)
	hashPasword, err := HashPassword(password)
	if err != nil {
		// Handle error
	}
	// ... extract other details as needed

	// Now, generate a JWT for the user
	jwtToken, err := GenerateJWT(/* user data, if needed */)
	if err != nil {
		return c.Status(500).SendString("Failed to generate JWT: " + err.Error())
	}

	// Set the JWT in a cookie or return it in the response, based on your app's design
	c.Set("Authorization", "Bearer "+jwtToken)
	return c.Redirect("/api/survey")
}

func IsAuthenticated(c *fiber.Ctx) error {
	// Check if the user has an authenticated session or cookie
	// If yes, proceed
	// If not, redirect to login
}

