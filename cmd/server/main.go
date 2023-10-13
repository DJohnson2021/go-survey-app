package main

import (
	// "github.com/DJohnson2021/go-survey-app/db"
	"log"
	// "os"
	// "fmt"

	"github.com/DJohnson2021/go-survey-app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/DJohnson2021/go-survey-app/api/middleware"
)


func main() {
	utils.LoadEnv()
	middleware.InitOauthConfig()

	/*
	client_ID := os.Getenv("GOOGLE_CLIENT_ID")
	client_Secret := os.Getenv("GOOGLE_CLIENT_SECRET")
	fmt.Println("Client_ID: ", client_ID)
	fmt.Println("OauthConfig.ClientID: ", middleware.OauthConfig.ClientID)
	fmt.Println("Client_Secret: ", client_Secret)
	fmt.Println("OauthConfig.client_secret: ", middleware.OauthConfig.ClientSecret)
	*/


	app := fiber.New()

	// Home Route
	app.Get("/", HomePage)
	// OAuth Routes
	app.Get("/api/user/oauth2/google/login", middleware.OauthGoogleLogin)
	app.Get("/api/user/oauth2/google/callback", middleware.OauthGoogleCallBack)

	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
	
}

func HomePage(c *fiber.Ctx) error {
	c.Type("html") // Set Content-Type header to text/html; charset=utf-8
	htmlContent := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta http-equiv="X-UA-Compatible" content="IE=edge">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Login with Google</title>
		</head>
		<body>
			<a href="/api/user/oauth2/google/login">
				<button>Login with Google</button>
			</a>
		</body>
		</html>
	`
	return c.SendString(htmlContent)
}
