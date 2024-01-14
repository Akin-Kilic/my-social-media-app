package routes

import (
	"context"
	"errors"
	"net/http"
	"social-media-app/pkg/domains/post"
	"social-media-app/pkg/dtos"
	"social-media-app/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func PostRoutes(app *fiber.App, service post.Service) {

	app.Get("/post/sh/:shortLink", GetWithShortLink(service))
	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: jwtware.SigningKey{Key: []byte(config.ReadValue().JwtSecret)},
	// }))
	post := app.Group("/post")

	post.Get("/userposts", GetUserAllPosts(service))
	post.Get("/posts", GetAllPosts(service))
	post.Post("/post", CreatePost(service))
	post.Put("/post", UpdatePost(service))
	post.Delete("/post", DeletePost(service))
}

func GetUserAllPosts(s post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("Authorization")

		jc, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token errror")
		}

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		posts, err := s.GetUserAllPosts(context.Background(), jc.UserId)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get posts",
			})
		}
		return c.JSON(posts)
	}
}

func GetAllPosts(s post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		posts, err := s.GetAllPosts(context.Background())
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get posts",
			})
		}
		return c.JSON(posts)
	}
}

func GetPostWithId(s post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		postId := c.Params("postId")

		userUuid, err := uuid.Parse(userId)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}
		postUuid, err := uuid.Parse(postId)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		posts, err := s.GetPostsWithPostId(context.Background(), userUuid, postUuid)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get posts",
			})
		}
		return c.JSON(posts)
	}
}

func CreatePost(s post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req *dtos.AddPostDTO
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		token := c.Cookies("Authorization")

		jc, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token error")
		}

		if err := s.CreatePost(c.Context(), req, jc.UserId); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create post",
			})
		}

		return c.SendStatus(http.StatusCreated)
	}
}

func UpdatePost(s post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		token := c.Cookies("Authorization")

		js, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token failed")
		}

		var req dtos.UpdatePostDTO
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		if err := s.UpdatePost(c.Context(), &req, js.UserId); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update post",
			})
		}

		return c.SendStatus(http.StatusOK)
	}
}

func DeletePost(s post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		token := c.Cookies("Authorization")

		jc, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token failed")
		}

		var requestBody map[string]interface{}

		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		postId, ok := requestBody["post_id"].(string)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "post_id field is missing or not a number",
			})
		}

		postUuid, err := uuid.Parse(postId)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid post ID",
			})
		}

		if err := s.DeletePost(c.Context(), jc.UserId, postUuid); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete post",
			})
		}

		return c.SendStatus(http.StatusOK)
	}
}

func GetWithShortLink(s post.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		shortLink := c.Params("shortLink")

		post, err := s.GetWithShortLink(c.Context(), shortLink)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update post",
			})
		}
		return c.Status(200).JSON(post)
	}
}
