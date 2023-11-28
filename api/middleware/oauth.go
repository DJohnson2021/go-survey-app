package middleware

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/DJohnson2021/go-survey-app/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

var OauthConfig *oauth2.Config

// Parse the user data
var googleUser struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Email      string `json:"email"`
}

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

	if err := json.Unmarshal(data, &googleUser); err != nil {
		log.Println("Error unmarshaling user data:", err)
		return c.Redirect("/")
	}

	user, err := models.GetUserByID(googleUser.ID)
	if err != nil {
		log.Println("Database error, Error getting user by ID:", err)
		return c.Redirect("/")
	}

	if user == nil {
		// Create user
		newUser := &models.User{
			GoogleID:   googleUser.ID,
			Username:   googleUser.Name,
			GivenName:  googleUser.GivenName,
			FamilyName: googleUser.FamilyName,
			Email:      googleUser.Email,
			Created_At: time.Now(),
		}
		if err := models.CreateUser(newUser); err != nil {
			log.Println("Database error, Failed to create user:", err)
			return c.Redirect("/")
		}

		user = newUser
	}

	if user.Username == "" || user.Email == "" {
		log.Println("Error getting user name and email from database")
		return c.Redirect("/")
	}

	// Generate a JWT token for the newly registered user
	token, err := GenerateJWT(user.Username, user.Email)
	if err != nil {
		log.Println("Error generating JWT token:", err)
		return c.Redirect("/")
	}

	// Set Secure flag to true for production
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		HTTPOnly: true,
		SameSite: "Lax",
		Expires:  time.Now().Add(24 * time.Hour), // e.g., expires in 24 hours
	})	

	return c.Redirect("/api/user/survey")
}

func generateStateOauthCookies(c *fiber.Ctx) string {
	var expiration = time.Now().Add(20 * time.Minute)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	fiberCookie := &fiber.Cookie{
		Name:     "oauthstate",
		Value:    state,
		Expires:  expiration,
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
