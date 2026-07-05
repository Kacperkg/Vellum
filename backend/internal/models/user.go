package models

type User struct {
	BaseModel
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`

	Email        string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
}
