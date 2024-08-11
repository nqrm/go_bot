package service

import (
	"context"

	model "github.com/nqrm/go_bot/internal/model/expenses"
)

type ExpenseService interface {
	Create(ctx context.Context, info *model.ExpenseInfo) (string, error)
	Get(ctx context.Context, uuid string) (*model.Expense, error)
}
