package main

import (
	"context"
	"fmt"
	"log"
	"social-media-app/internal/api/routes"
	"social-media-app/pkg/config"
	"social-media-app/pkg/db"
	"social-media-app/pkg/domains/comment"
	"social-media-app/pkg/domains/friends"
	"social-media-app/pkg/domains/like"
	"social-media-app/pkg/domains/post"
	"social-media-app/pkg/domains/user"
	"social-media-app/pkg/redis"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	PrometheusUserRegistrationTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "social_media_app_user_registration_total",
			Help: "Total number of user registrations",
		},
		[]string{},
	)

	PrometheusUserLoginTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "social_media_app_user_login_total",
			Help: "Total number of user logins",
		},
		[]string{},
	)
)

func init() {
	prometheus.MustRegister(
		PrometheusUserRegistrationTotal,
		PrometheusUserLoginTotal,
	)
}

func main() {
	conf := config.ReadValue()
	db.Connect(
		conf.Database,
	)

	redis.Connect(conf.Redis)
	err2 := redis.Ping(context.Background())
	if err2 != nil {
		fmt.Println("Error connecting to redis from ping")
	}

	app := fiber.New()
	db := db.Client()
	userRepo := user.NewRepository(db)
	userService := user.NewUser(userRepo)

	repo := post.NewRepository(db)
	postService := post.NewPost(repo)

	likeRepo := like.NewRepository(db)
	likeService := like.NewLike(likeRepo)

	commentRepo := comment.NewRepository(db)
	commetService := comment.NewComment(commentRepo)

	friendRepo := friends.NewRepository(db)
	friendService := friends.NewFriends(friendRepo)

	routes.UserRoutes(app, userService)
	routes.PostRoutes(app, postService)
	routes.LikeRoutes(app, likeService)
	routes.CommentRoutes(app, commetService)
	routes.FriendRoutes(app, friendService)

	err := app.Listen(":" + config.ReadValue().Port)
	if err != nil {
		log.Fatal(err)
	}
}
