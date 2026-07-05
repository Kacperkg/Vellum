package dto

import (
	"github.com/google/uuid"
	"github.com/kacperkg/vellum/internal/models"
)

type CreateBudgetRequest struct {
	CategoryID uuid.UUID

	AllocatedAmount float64
}

type UpdateBudgetRequest struct {
	CategoryID uuid.UUID

	AllocatedAmount float64
}

type BudgetResponse struct {
	ID         uuid.UUID
	CategoryID uuid.UUID
	Category   models.Category

	AllocatedAmount float64
}
