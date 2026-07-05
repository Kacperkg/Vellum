package models

import "github.com/google/uuid"

type Goal struct {
	BaseModel

	UserID uuid.UUID `gorm:"type:uuid;not null"`
	User   User      `gorm:"foreignKey:UserID"`

	GoalName            string  `gorm:"not null"`
	TargetAmount        float64 `gorm:"not null"`
	SavedAmount         float64 `gorm:"not null"`
	MonthlyContribution float64 `gorm:"not null"`
}
