package repositories

import (
	"github.com/google/uuid"
	"github.com/kacperkg/vellum/internal/models"
)

type CategoryRepository interface {
	FindByID(id uuid.UUID) (*models.Category, error)
	FindAll() ([]*models.Category, error)
}