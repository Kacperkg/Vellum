package models

import "github.com/google/uuid"

type Account struct {
	BaseModel

	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User   User      `gorm:"foreignKey:UserID"`

	Name string `gorm:"not null"`
}
