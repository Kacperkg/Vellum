package models

type User struct {
	BaseModel
	FirstName string `gorm:"not null"`
	LastName string `gorm:"not null"`
	Currency string `gorm:"not null;default:'GBP'"`

	Email string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
}