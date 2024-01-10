package main

import (
	"log"
	"social-media-app/internal/api/routes"
	"social-media-app/pkg/config"
	"social-media-app/pkg/db"
	"social-media-app/pkg/domains/user"
	"social-media-app/pkg/redis"

	"github.com/gofiber/fiber/v2"
)

func main() {
	conf := config.ReadValue()
	db.Connect(
		conf.Database,
	)

	redis.Connect(conf.Redis)

	app := fiber.New()
	db := db.Client()
	userRepo := user.NewRepository(db)
	userService := user.NewUser(userRepo)

	routes.UserRoutes(app, userService)

	err := app.Listen(":" + config.ReadValue().Port)
	if err != nil {
		log.Fatal(err)
	}
}
