package controllers

import (
	"fmt"
	"github.com/DJohnson2021/go-survey-app/api/middleware"
	"github.com/DJohnson2021/go-survey-app/db"
	"github.com/DJohnson2021/go-survey-app/models"
	"github.com/DJohnson2021/go-survey-app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"time"
	//"strconv"
	// "log"
	// "html/template"
)

// HomePage is the controller for the home page.
func HomePage(c *fiber.Ctx) error {
	return c.SendFile("../../templates/Homepage.html")
}

// LoginPage is the controller for the login page.
func LoginPage(c *fiber.Ctx) error {
	return c.SendFile("../../templates/Login.html")
}

// SurveyPage is the controller for the survey page.
func SurveyPage(c *fiber.Ctx) error {
	return c.SendFile("../../templates/survey.html")
}

// ResultPage is the controller for the survey results page.
func ResultPage(c *fiber.Ctx) error {
	return c.SendFile("../../templates/results.html")
}

func SubmitSurvey(c *fiber.Ctx) error {
	// Extract the JWT token from the cookie
	tokenString := c.Cookies("jwt")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	// Decode the JWT token to get user data
	claims := &middleware.Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		jwtKey, err := utils.GetJWTSecret()
		if err != nil {
			return nil, fmt.Errorf("error getting JWT secret: %v", err)
		}
		return jwtKey, nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
	}

	// Get user ID from the database based on email
	user, err := models.GetUserByEmail(claims.Email)
	if err != nil {
		// Handle error (e.g., user not found)
        return fmt.Errorf("error find user with this email: %v", claims.Email)
	}
	userID := user.ID

	// Parse form values
	sleepAmount := c.FormValue("sleep_amount")
	friendSurvey := c.FormValue("friend_survey")

	// Convert questionID to int32
	/*
	   questionIDInt, err := strconv.ParseInt(questionID, 10, 32)
	   if err != nil {
	       // handle error
	   }
	*/

	// Create a new response
	responseQ1 := models.Response{
		Question_id:   1,
		User_id:       userID,
		Response: sleepAmount,
		Created_At:    time.Now(),
	}

	responseQ2 := models.Response{
		Question_id:   2,
		User_id:       userID,
		Response: friendSurvey,
		Created_At:    time.Now(),
	}

	// Save the response to the database
	if err := db.DB.Create(&responseQ1).Error; err != nil {
		return fmt.Errorf("error creating user response to question 1: %v", err)
	}

	if err := db.DB.Create(&responseQ2).Error; err != nil {
		return fmt.Errorf("error creating user response to question 2: %v", err)
	}

	return c.Redirect("/api/user/survey/results")
}
