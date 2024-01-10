package db

import (
	"fmt"
	"log"
	"social-media-app/pkg/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBClient *gorm.DB

func Connect(db config.Database) {
	if DBClient != nil {
		return
	}
	var err error
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", db.Host, db.Port, db.User, db.Password, db.Name)
	DBClient, err = gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  dsn,
				PreferSimpleProtocol: true,
			},
		))
	if err != nil {
		log.Println("gorm open error:", err)
	}
	sqlDB, err := DBClient.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	sqlDB.SetConnMaxIdleTime(time.Minute * 5)

	if err != nil {
		panic(err)
	}
	if db.Migrate {
		AutoMigrate()
	}
}

func Client() *gorm.DB {
	return DBClient
}
