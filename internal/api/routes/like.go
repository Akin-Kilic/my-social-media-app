package routes

import (
	"errors"
	"fmt"
	"net/http"
	"social-media-app/pkg/domains/like"
	"social-media-app/pkg/model"
	"social-media-app/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func LikeRoutes(app *fiber.App, service like.Service) {
	like := app.Group("/like")

	like.Post("/like", Create(service))
	like.Get("postlikes", GetLikesForPost(service))
	like.Get("commentlikes", GetLikesForComment(service))
	like.Get("/likecount", GetLikesCountForPost(service))
	like.Delete("/like", DeleteLike(service))

}

func Create(s like.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("handler başladı")

		token := c.Cookies("Authorization")

		jc, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token error")
		}

		var like model.Likes

		if err := c.BodyParser(&like); err != nil {
			fmt.Println("parse hatası")
			return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON format"})
		}
		fmt.Println("parse tamamlandı")

		fmt.Println("user id: ", jc.UserId)
		like.UserId = jc.UserId

		fmt.Println(like)
		if err := s.Create(c.Context(), &like); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to get like"})
		}
		return c.Status(201).JSON(fiber.Map{"message": "like create successfully"})
	}
}

func GetLikesForPost(s like.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("Authorization")
		_, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token error")
		}
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		var requestBody map[string]interface{}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// post_id'yi almak
		postId, ok := requestBody["post_id"].(string)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "post_id field is missing or not a string",
			})
		}

		uuidPostId, err := uuid.Parse(postId)
		if err != nil {
			return errors.New("parse uuid error")
		}

		posts, err := s.GetLikesWithPostId(c.Context(), uuidPostId)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get posts",
			})
		}
		return c.Status(200).JSON(posts)
	}
}

func GetLikesForComment(s like.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("Authorization")
		_, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token error")
		}
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid user ID",
			})
		}

		var requestBody map[string]interface{}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// post_id'yi almak
		postId, ok := requestBody["comment_id"].(string)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "post_id field is missing or not a string",
			})
		}

		uuidCommentId, err := uuid.Parse(postId)
		if err != nil {
			return errors.New("parse uuid error")
		}

		posts, err := s.GetLikesWithCommentId(c.Context(), uuidCommentId)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to get posts",
			})
		}
		return c.Status(200).JSON(posts)
	}
}

func GetLikesCountForPost(s like.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var requestBody map[string]interface{}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// post_id'yi almak
		postId, ok := requestBody["post_id"].(string)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "post_id field is missing or not a string",
			})
		}
		fmt.Println("post id: ", postId)
		count, err := s.GetLikesCountForPost(c.Context(), postId)
		fmt.Println("count: ", count)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "cannot get likes count for post"})
		}
		return c.Status(200).JSON(fiber.Map{"message": "like get successfully", "count": count})
	}
}

func DeleteLike(s like.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {

		token := c.Cookies("Authorization")

		_, err := utils.ParseToken(token)
		if err != nil {
			return errors.New("parse token error")
		}

		var requestBody map[string]interface{}
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		// post_id'yi almak
		likeId, ok := requestBody["like_id"].(string)
		if !ok {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "post_id field is missing or not a string",
			})
		}

		err = s.DeleteLike(c.Context(), likeId)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "cannot delete likes"})
		}
		return c.Status(200).JSON(fiber.Map{"message": "like deleted successfully"})
	}
}
