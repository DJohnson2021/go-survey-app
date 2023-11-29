package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/DJohnson2021/go-survey-app/models"
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
    // Assuming you have user authentication and can get the user's ID
    userID := 

    // Parse form values
    question := "How much do you sleep?" // This can be dynamic based on your form
    answer := c.FormValue("answer")

    // Create a new survey response
    response := models.SurveyResponse{
        UserID:   userID,
        Question: question,
        Answer:   answer,
    }

    // Save the response to the database
    result := db.DB.Create(&response)
    if result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).SendString("Error saving response")
    }

    return c.Redirect("/survey-success")
}