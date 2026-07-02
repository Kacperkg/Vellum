package models

import (
	"github.com/google/uuid"
)

type UserSettings struct {
    BaseModel

    UserID uuid.UUID

    Currency string
    DateFormat string
    NumberFormat string
    WeekStartsOn string
}