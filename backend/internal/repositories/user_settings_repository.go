package repositories

import (
	"github.com/google/uuid"
	"github.com/kacperkg/vellum/internal/models"
)

type UserSettingsRepository interface {
	FindByUserID(userID uuid.UUID) (*models.UserSettings, error)
	Update(settings *models.UserSettings) error
}