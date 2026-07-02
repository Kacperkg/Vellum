package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	BaseModel

	Name string `gorm:"not null"`
	Amount float64 `gorm:"not null"`

	CategoryID uuid.UUID `gorm:"type:uuid;not null"`
	Category   Category  `gorm:"foreignKey:CategoryID"`

	Date time.Time `gorm:"not null"`
	
	Frequency Frequency      `gorm:"not null"`
	Type      TransactionType `gorm:"not null"`
}	