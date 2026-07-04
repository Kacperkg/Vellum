package repositories

import (
	"time"

	"github.com/google/uuid"
	"github.com/kacperkg/vellum/internal/models"
	"gorm.io/gorm"
)

type GormTransactionRepository struct {
	db *gorm.DB
}

func NewGormTransactionRepository(db *gorm.DB) *GormTransactionRepository {
	return &GormTransactionRepository{
		db: db,
	}
}

func (r *GormTransactionRepository) Create(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *GormTransactionRepository) FindByID(id uuid.UUID) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.First(&transaction, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

func (r *GormTransactionRepository) FindByUserID(userID uuid.UUID) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.Where("user_id = ?", userID).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *GormTransactionRepository) Update(transaction *models.Transaction) error {
	return r.db.Save(transaction).Error
}

func (r *GormTransactionRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Transaction{}, "id = ?", id).Error
}

func (r *GormTransactionRepository) FindByAccountID(accountID uuid.UUID) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.Where("account_id = ?", accountID).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *GormTransactionRepository) FindByCategoryID(userID, categoryID uuid.UUID) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.Where("category_id = ? AND user_id = ?", categoryID, userID).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *GormTransactionRepository) FindByDateRange(userID uuid.UUID, startDate, endDate time.Time) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.Where("date >= ? AND date <= ? AND user_id = ?", startDate, endDate, userID).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *GormTransactionRepository) FindByFrequency(userID uuid.UUID, frequency models.Frequency) ([]*models.Transaction, error) {
	var transactions []*models.Transaction
	err := r.db.Where("frequency = ? AND user_id = ?", frequency, userID).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return transactions, nil
}