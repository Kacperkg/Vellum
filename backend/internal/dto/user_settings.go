package dto

import (
	"github.com/google/uuid"
	"github.com/kacperkg/vellum/internal/models"
)

type UpdateUserSettingsRequest struct {
	Currency     models.Currency
	DateFormat   models.DateFormat
	NumberFormat models.NumberFormat
	WeekStartsOn models.WeekStartsOn
}

type UserSettingsResponse struct {
	ID           uuid.UUID
	Currency     models.Currency
	DateFormat   models.DateFormat
	NumberFormat models.NumberFormat
	WeekStartsOn models.WeekStartsOn
}
