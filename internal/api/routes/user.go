package routes

import (
	"context"
	"social-media-app/pkg/config"
	"social-media-app/pkg/domains/user"
	"social-media-app/pkg/dtos"
	"social-media-app/pkg/model"
	"social-media-app/pkg/utils"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, service user.Service) {
	app.Post("/register", Register(service))
	app.Post("/login", Login(service))

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.ReadValue().JwtSecret)},
	}))
	

}

func Register(s user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user *model.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload" + err.Error()})
		}
		// if v.HasErrors() {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Validation error", "errors": v.FieldErrors})
		// 	return
		// }
		ctx := context.Background()
		err := s.Register(ctx, user)
		if err != nil {
			return c.Status(404).JSON(utils.Response(err.Error()))
		}
		return c.Status(201).JSON(utils.Response(user))
	}
}

func Login(s user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			req dtos.LoginReq
		)
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload" + err.Error()})
		}
		// if v.HasErrors() {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Validation error", "errors": v.FieldErrors})
		// 	return
		// }

		ctx := context.Background()
		ipCtx := context.WithValue(ctx, "ip_address", c.IP())

		loginRes, err := s.Login(ipCtx, req)
		if err != nil {
			return c.Status(404).JSON(utils.Response(err.Error()))
		}
		return c.Status(201).JSON(utils.Response(loginRes))
	}
}
