package models

type Account struct {
	BaseModel

	Name string `gorm:"uniqueIndex;not null"`
}