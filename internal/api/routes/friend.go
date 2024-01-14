package routes

import (
	"errors"
	"fmt"
	"net/http"
	"social-media-app/pkg/domains/friends"
	"social-media-app/pkg/dtos"
	"social-media-app/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func FriendRoutes(app *fiber.App, service friends.Service) {
	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte(config.ReadValue().JwtSecret)},
	// }))
	friend := app.Group("/friend")
	friend.Post("friend", AddFriend(service))
	friend.Get("friends", GetFriends(service))
	friend.Put("accept-friend", AcceptFriend(service))
	friend.Put("reject-friend", RejectFriend(service))
	friend.Delete("friend", DeleteFriend(service))
}

func AddFriend(s friends.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		token := c.Cookies("Authorization")

		jc, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token errror")
		}

		var addFriendDTO dtos.AddFriendDTO

		err = c.BodyParser(&addFriendDTO)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to body parse",
			})
		}

		err = s.AddFriend(c.Context(), &addFriendDTO, jc.UserId)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get posts",
			})
		}
		fmt.Println("control çıktı")
		return c.Status(201).JSON(fiber.Map{
			"message": "Friend created successfully",
		})
	}
}

func GetFriends(s friends.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		token := c.Cookies("Authorization")

		jc, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token errror")
		}

		friends, err := s.GetFriends(c.Context(), jc.UserId)
		fmt.Println("çıktı")
		fmt.Println("çıktı: ", friends)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get friends",
			})
		}
		return c.Status(201).JSON(fiber.Map{
			"friends": friends,
		})
	}
}

func AcceptFriend(s friends.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		token := c.Cookies("Authorization")

		jc, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token errror")
		}

		var requestBody map[string]interface{}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		friendId, ok := requestBody["friend_id"].(string)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "post_id field is missing or not a string",
			})
		}

		uuuidFriendId, err := uuid.Parse(friendId)
		if err != nil {
			return errors.New("uuid parse error")
		}

		err = s.AcceptFriend(c.Context(), jc.UserId, uuuidFriendId)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to accept friend",
			})
		}
		return c.Status(201).JSON(fiber.Map{
			"message": "Friend accepted",
		})
	}
}

func RejectFriend(s friends.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		token := c.Cookies("Authorization")

		jc, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token errror")
		}

		var requestBody map[string]interface{}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		friendId, ok := requestBody["friend_id"].(string)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "post_id field is missing or not a string",
			})
		}

		uuuidFriendId, err := uuid.Parse(friendId)
		if err != nil {
			return errors.New("uuid parse error")
		}

		err = s.RejectFriend(c.Context(), jc.UserId, uuuidFriendId)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to reject friend",
			})
		}
		return c.Status(201).JSON(fiber.Map{
			"message": "Friend rejected",
		})
	}
}

func DeleteFriend(s friends.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		token := c.Cookies("Authorization")

		jc, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token errror")
		}

		var requestBody map[string]interface{}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		friendId, ok := requestBody["friend_id"].(string)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "post_id field is missing or not a string",
			})
		}

		uuuidFriendId, err := uuid.Parse(friendId)
		if err != nil {
			return errors.New("uuid parse error")
		}

		err = s.DeleteFriend(c.Context(), jc.UserId, uuuidFriendId)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete friend",
			})
		}
		return c.Status(201).JSON(fiber.Map{
			"message": "Friend deleted",
		})
	}
}
