package middleware

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/DJohnson2021/go-survey-app/models"
)

/*
func LoadEnv() {
	// Load .env file
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}
*/


var OauthConfig *oauth2.Config

func InitOauthConfig() {
	OauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/api/user/oauth2/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func OauthGoogleLogin(c *fiber.Ctx) error {
	oauthState := generateStateOauthCookies(c)
	u := OauthConfig.AuthCodeURL(oauthState)
	return c.Redirect(u)
}

func OauthGoogleCallBack(c *fiber.Ctx) error {
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
	 // Parse the user data
	 var googleUser struct {
        ID       string `json:"id"`
        Name 	 string `json:"name"`
        Email    string `json:"email"`
        // ... any other fields you want to capture
    }

	if err := json.Unmarshal(data, &googleUser); err != nil {
		log.Println("Error unmarshaling user data:", err)
		return c.Redirect("/")
	}

	user, err := models.GetUserByID(googleUser.ID)
	if err != nil {
		log.Println("Database error, Error getting user by ID:", err)
		return c.Redirect("/")
	}

	if user != nil {
		// Generate a JWT token for registered user
		// token, err := middleware.GenerateJWT(/* user data, if needed */)
		return c.Redirect("/api/user/profile")
	}


	if user == nil {
		google_ID, err := strconv.ParseInt(googleUser.ID, 10, 64)
		if err != nil {
			log.Println("Error converting googleUser.ID to int64:", err)
			return c.Redirect("/")
		}

		// Create user
		newUser := &models.User{
			GoogleID: google_ID,
			Username: googleUser.Name,
			Email:    googleUser.Email,
			Timestamp: time.Now(),
		}
		if err := models.CreateUser(newUser); err != nil {
			log.Println("Database error, Failed to create user:", err)
			return c.Redirect("/")
		}

		// Generate a JWT token for the newly registered user
		// token, err := middleware.GenerateJWT(/* user data, if needed */)
	}

	// Redirect or response with a token.
	return c.Redirect("/api/user/profile")


	// More code .....

	// return c.SendString(fmt.Sprintf("UserInfo: %s\n", data))
}

func generateStateOauthCookies(c *fiber.Ctx) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	fiberCookie := &fiber.Cookie{
		Name:    "oauthstate",
		Value:   state,
		Expires: expiration,
		HTTPOnly: true,
	}

	c.Cookie(fiberCookie)

	return state
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	token, err := OauthConfig.Exchange(context.Background(), code)
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
