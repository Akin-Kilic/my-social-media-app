package db

import "social-media-app/pkg/model"

func AutoMigrate() {
	DBClient.AutoMigrate(
		&model.User{},
	)

}
