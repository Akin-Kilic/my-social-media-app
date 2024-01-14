package db

import "social-media-app/pkg/model"

func AutoMigrate() {
	DBClient.AutoMigrate(
		&model.User{},
		&model.Comment{},
		&model.Friend{},
		&model.Image{},
		&model.Likes{},
		&model.Post{},
	)
}
