package models

import "github.com/google/uuid"

type Budget struct {
	BaseModel

	CategoryID uuid.UUID `gorm:"type:uuid;not null"`
	Category Category  `gorm:"foreignKey:CategoryID"`

	AllocatedAmount float64 `gorm:"not null"`
}