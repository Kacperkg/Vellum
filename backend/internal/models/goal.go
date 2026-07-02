package models

type Goal struct {
	BaseModel
	
	GoalName string `gorm:"not null"`
	TargetAmount float64 `gorm:"not null"`
	SavedAmount float64 `gorm:"not null"`
	MonthlyContribution float64 `gorm:"not null"`
}