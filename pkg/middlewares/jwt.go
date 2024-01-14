package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func SetToken() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		// Check if the header is missing or doesn't have the correct format
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Extract the token from the header
		jwtToken := strings.TrimPrefix(authHeader, "Bearer ")

		// Now, 'jwtToken' contains the JWT token
		c.Locals("jwtToken", jwtToken)

		// Continue to the next middleware or route handler
		return c.Next()
	}

}
