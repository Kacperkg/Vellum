package repositories

import (
	"errors"

	"github.com/google/uuid"
	"github.com/kacperkg/vellum/internal/models"
	"gorm.io/gorm"
)

type GormUserSettingsRepository struct {
	db *gorm.DB
}

func NewGormUserSettingsRepository(db *gorm.DB) *GormUserSettingsRepository {
	return &GormUserSettingsRepository{
		db: db,
	}
}

func (r *GormUserSettingsRepository) FindByUserID(userID uuid.UUID) (*models.UserSettings, error) {
	var settings models.UserSettings
	err := r.db.First(&settings, "user_id = ?", userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &settings, nil
}

func (r *GormUserSettingsRepository) Update(settings *models.UserSettings) error {
	return r.db.Save(settings).Error
}

func (r *GormUserSettingsRepository) Create(settings *models.UserSettings) error {
	return r.db.Create(settings).Error
}
