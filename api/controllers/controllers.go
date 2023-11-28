package controllers

import (
	"github.com/gofiber/fiber/v2"
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