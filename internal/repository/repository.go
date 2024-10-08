package repository

import (
	"context"

	model "github.com/nqrm/go_bot/internal/model/expenses"
)

type ExpenseRepository interface {
	Create(ctx context.Context, uuid string, info *model.ExpenseInfo) error
	Get(ctx context.Context, uuid string) (*model.Expense, error)
}
