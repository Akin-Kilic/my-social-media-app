package routes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"social-media-app/pkg/constant"
	"social-media-app/pkg/domains/user"
	"social-media-app/pkg/dtos"
	"social-media-app/pkg/model"
	"social-media-app/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UserRoutes(app *fiber.App, service user.Service) {
	user := app.Group("/user")
	user.Post("/register", Register(service))
	user.Post("/login", Login(service))
	user.Post("/logout", Logout(service))
	user.Post("/image", UploadImage(service))
	user.Get("/profile", GetProfilePhoto(service))
}

func Register(s user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user model.User
		if err := c.BodyParser(&user); err != nil {
			log.Println("Error parsing request body:", err)
			return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON format"})
		}

		if err := s.Register(c.Context(), &user); err != nil {
			log.Println("Error registering user:", err)
			return c.Status(500).JSON(fiber.Map{"error": "Failed to register user"})
		}

		return c.Status(201).JSON(fiber.Map{"message": "User registered successfully"})
	}
}

func Login(s user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user dtos.LoginReq
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON format"})
		}

		loginDto, err := s.Login(c.Context(), &user)

		if err != nil {
			log.Println("Error during login:", err)
			return c.Status(500).JSON(fiber.Map{"error": "Failed to log in user"})
		}

		c.Cookie(&fiber.Cookie{
			Name:     "Authorization",
			Value:    loginDto.Token,
			Expires:  time.Now().Add(time.Minute * 5),
			SameSite: "Lax",
			HTTPOnly: true,
		})

		return c.Status(200).JSON(fiber.Map{
			"token": loginDto.Token,
		})
	}
}

func Logout(s user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error
		token := c.Cookies("Authorization")

		jc, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token errror")
		}

		key := fmt.Sprintf(constant.RedisForJwt, token, jc.UserId)

		ctx := context.Background()
		err = s.Logout(ctx, key)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to register user"})
		}

		c.Cookie(&fiber.Cookie{
			Name:     "Authorization",
			Value:    "",
			SameSite: "Lax",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		})

		return c.Status(201).JSON(fiber.Map{
			"message": "User logout successfully",
		})
	}
}

func UploadImage(s user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("Authorization")

		jc, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token error: ")
		}
		file, err := c.FormFile("image")
		if err != nil {
			return err
		}
		path := "./images/" + file.Filename
		if err := c.SaveFile(file, path); err != nil {
			return err
		}
		imageType := c.FormValue("image_type")
		userId := jc.UserId
		entityId := c.FormValue("entityId")

		entityUuid, err := uuid.Parse(entityId)
		if err != nil {
			return err
		}
		if err := s.SaveImage(path, imageType, userId, entityUuid); err != nil {
			return err
		}
		return c.SendStatus(http.StatusOK)
	}
}

func GetProfilePhoto(s user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("Authorization")
		jc, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token error")
		}
		image, err := s.GetProfilePhoto(c.Context(), jc.UserId)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": err.Error()})
		}
		byteImage, err := json.Marshal(image)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(200).JSON(fiber.Map{
			"message": "Profile photo get successfully",
			"data":    byteImage,
		})
	}
}

// func RegisterPrometheusRoute(a *fiber.App) {
// 	a.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))
// }
