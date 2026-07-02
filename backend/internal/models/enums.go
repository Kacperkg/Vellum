package models

type Frequency string

const (
	OneOff  Frequency = "one-off"
	Daily   Frequency = "daily"
	Weekly  Frequency = "weekly"
	Monthly Frequency = "monthly"
	Yearly  Frequency = "yearly"
)

type TransactionType string

const (
	Income TransactionType = "income"
	Expense TransactionType = "expense"
)