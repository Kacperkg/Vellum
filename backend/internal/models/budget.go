package models

import "github.com/google/uuid"

type Budget struct {
	BaseModel

	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User   User      `gorm:"foreignKey:UserID"`

	CategoryID uuid.UUID `gorm:"type:uuid;not null"`

	AllocatedAmount float64 `gorm:"not null"`
}
