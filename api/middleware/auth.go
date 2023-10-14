package middleware

import (
	//"context"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	//"github.com/golang-jwt/jwt"
)

func IsAdminAuthorized() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract token from the Authorization header
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			log.Println("Error authorizing admin: empty token string")
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing authorization token",
			})
			return c.Redirect("/")
		}

		// Remove 'Bearer ' prefix from tokenString
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Use the adminAuthorized function to verify the token
		claims, err := adminAuthorized(tokenString)
		if err != nil {
			log.Println("Error authorizing user: ", err)
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "You are not authorized to access this resource",
			})
			return c.Redirect("/")
		}

		// You can also add further checks here if needed
		if !claims.isAdmin {
			log.Println("Error authorizing user: ", err)
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "You are not authorized to access this resource",
			})
			return c.Redirect("/")
		}
		return c.Next()
	}
}

func IsUserAuthorized() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract token from the Authorization header
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			log.Println("Error authorizing user: empty token string")
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing authorization token",
			})
			return c.Redirect("/")
		}

		// Remove 'Bearer ' prefix from tokenString if it exists
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// Use the userAuthorized function to verify the token
		_, err := userAuthorized(tokenString)
		if err != nil {
			log.Println("Error authorizing user: ", err)
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "You are not authorized to access this resource",
			})
			return c.Redirect("/")
		}

		// Checking claims.isAdmin is not necessary here because and admin can access all user routes

		return c.Next()
	}
}
