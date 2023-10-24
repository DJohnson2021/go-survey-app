package main

import (
	//"fmt"
	"log"

	"github.com/DJohnson2021/go-survey-app/db"
	"github.com/DJohnson2021/go-survey-app/api/controllers"

	//"os"

	"github.com/DJohnson2021/go-survey-app/api/middleware"
	"github.com/DJohnson2021/go-survey-app/utils"
	"github.com/gofiber/fiber/v2"
)


func main() {
	utils.LoadEnv()
	middleware.InitOauthConfig()
	db.InitDatabase()
	defer db.CloseDatabase()

	app := fiber.New()

	// Home Route
	app.Static("/", "../../templates/")
	app.Static("/login", "../../templates/")
	app.Static("/static", "../../static")
	app.Get("/", controllers.HomePage)
	// OAuth Routes
	// app.Static("/login", "../../templates/Login.html")
	app.Get("/login", controllers.LoginPage)
	app.Get("/api/user/oauth2/google/login", middleware.OauthGoogleLogin)
	app.Get("/api/user/oauth2/google/callback", middleware.OauthGoogleCallBack)
	
	// Test routes
	app.Get("/api/user/dashboard", middleware.IsUserAuthorized(), func(c *fiber.Ctx) error {	
		c.Type("html")
		html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Admin Check</title>
		</head>
		<body>
			<h2>Hello user!</h2>
			<p>Are you an admin?</p>
			<form action="/api/admin/dashboard" method="GET">
				<button type="submit">I am admin</button>
			</form>
		</body>
		</html>
		`
		return c.SendString(html)
	})	

	app.Get("/api/admin/dashboard",middleware.IsAdminAuthorized(), func(c *fiber.Ctx) error {
		return c.SendString("Hello, admin!")
	})

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

