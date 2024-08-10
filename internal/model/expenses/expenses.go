package model

import "time"

type Expenses struct {
	UUID      string
	CreatedAt time.Time
	Category  string
	Amount    int
}

type ExpensesInfo struct {
	CreatedAt time.Time
	Category  string
	Amount    int
}
