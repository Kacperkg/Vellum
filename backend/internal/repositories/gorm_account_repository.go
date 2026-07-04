package repositories

import (
	"github.com/google/uuid"
	"github.com/kacperkg/vellum/internal/models"
	"gorm.io/gorm"
)

type GormAccountRepository struct {
	db *gorm.DB
}

func NewGormAccountRepository(db *gorm.DB) *GormAccountRepository {
	return &GormAccountRepository{
		db: db,
	}
}

func (r *GormAccountRepository) Create(account *models.Account) error {
	return r.db.Create(account).Error
}

func (r *GormAccountRepository) GetByID(id uuid.UUID) (*models.Account, error) {
	var account models.Account
	err := r.db.First(&account, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *GormAccountRepository) FindByUserID(userID uuid.UUID) ([]*models.Account, error) {
	var accounts []*models.Account
	err := r.db.Where("user_id = ?", userID).Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (r *GormAccountRepository) Update(account *models.Account) error {
	return r.db.Save(account).Error
}

func (r *GormAccountRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Account{}, "id = ?", id).Error
}
