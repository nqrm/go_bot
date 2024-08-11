package model

import "time"

type Expense struct {
	UUID      string
	CreatedAt time.Time
	Category  string
	Amount    int
}

type ExpenseInfo struct {
	CreatedAt time.Time
	Category  string
	Amount    int
}
