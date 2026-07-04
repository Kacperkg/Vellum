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

type NumberFormat string

const (
	CommaDecimal NumberFormat = "comma_decimal" // 1,234.56
	DotDecimal   NumberFormat = "dot_decimal"   // 1.234,56
	Apostrophe   NumberFormat = "apostrophe"    // 1'234.56
)