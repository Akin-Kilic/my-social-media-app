package middlewares

import (
	"social-media-app/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var secret = config.ReadValue().JwtSecret

func JwtControl() fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		jwtToken, err := jwt.ParseWithClaims(token, &customClaims, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})

		// Devam et
		return c.Next()
	}
}
