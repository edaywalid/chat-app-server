package db

import (
	"github.com/edaywalid/chat-app/configs"
	"github.com/edaywalid/chat-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *configs.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DBUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
