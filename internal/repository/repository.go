package repository

import (
	"context"

	model "github.com/nqrm/go_bot/internal/model/expenses"
)

type ExpensesRepository interface {
	Create(ctx context.Context, uuid string, info *model.ExpensesInfo) error
	Get(ctx context.Context, uuid string) (*model.ExpensesInfo, error)
}
