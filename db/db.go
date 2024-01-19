package db

import (
	"os"

	"example.com/poll-app-backend-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDB() error {

	var err error

	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	err = DB.AutoMigrate(&models.User{}, &models.Poll{}, &models.Category{}, &models.Option{}, &models.Vote{})

	return err
}
