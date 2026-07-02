package database

import (
	"github.com/kacperkg/vellum/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
	)
}