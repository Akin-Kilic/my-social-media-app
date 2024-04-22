package routes

import (
	"errors"
	"fmt"
	"net/http"
	"social-media-app/pkg/domains/comment"
	"social-media-app/pkg/dtos"
	"social-media-app/pkg/model"
	"social-media-app/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CommentRoutes(app *fiber.App, service comment.Service) {
	comment := app.Group("/comment")
	comment.Post("comment", AddComment(service))
	comment.Get("comments", GetCommentsForPost(service))
	comment.Put("comment", UpdateComment(service))
	comment.Delete("comment", DeleteComment(service))
}

func AddComment(s comment.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("Authorization")
		jc, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token error")
		}
		var comment model.Comment
		if err := c.BodyParser(&comment); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON format"})
		}
		comment.UserID = jc.UserId
		if err := s.AddComment(c.Context(), &comment); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to get comment"})
		}
		return c.Status(201).JSON(fiber.Map{"message": "comment create successfully"})
	}
}

func GetCommentsForPost(s comment.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody map[string]interface{}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		postId, ok := requestBody["post_id"].(string)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "post_id field is missing or not a string",
			})
		}

		postUuid, err := uuid.Parse(postId)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}
		posts, err := s.GetCommentsForPost(c.Context(), postUuid)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}
		return c.Status(200).JSON(fiber.Map{
			"data": posts,
		})
	}
}

func UpdateComment(s comment.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("Authorization")

		_, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token error")
		}

		var commentDto dtos.UpdateCommentDTO

		if err := c.BodyParser(&commentDto); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Failed to parse comment"})
		}

		fmt.Println(commentDto)

		err = s.UpdateComment(c.Context(), &commentDto)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed update comment"})
		}
		return c.Status(200).JSON(fiber.Map{"message": "Update comment successfully"})
	}
}

func DeleteComment(s comment.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody map[string]interface{}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		commentId, ok := requestBody["comment_id"].(string)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "post_id field is missing or not a string",
			})
		}

		commentUuid, err := uuid.Parse(commentId)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid userId"})
		}

		err = s.DeletePost(c.Context(), commentUuid)
		if err != nil {
			return c.Status(200).JSON(fiber.Map{"error": "Failed delete comment"})
		}
		return c.Status(200).JSON(fiber.Map{"message": "Delete comment successfully"})
	}
}
