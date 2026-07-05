package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/kacperkg/vellum/internal/models"
)

type CreateTransactionRequest struct {
	Name   string
	Amount float64

	AccountID uuid.UUID

	CategoryID uuid.UUID

	Date time.Time

	Frequency models.Frequency
	Type      models.TransactionType
}

type UpdateTransactionRequest struct {
	Name   string
	Amount float64

	AccountID uuid.UUID

	CategoryID uuid.UUID

	Date time.Time

	Frequency models.Frequency
	Type      models.TransactionType
}

type TransactionResponse struct {
	ID     uuid.UUID
	Name   string
	Amount float64

	AcountID uuid.UUID

	CategoryID uuid.UUID

	Date time.Time

	Frequency models.Frequency
	Type      models.TransactionType
}
