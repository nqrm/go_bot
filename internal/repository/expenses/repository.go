package expenses

import (
	"context"
	"sync"

	model "github.com/nqrm/go_bot/internal/model/expenses"
	def "github.com/nqrm/go_bot/internal/repository"
	repoModel "github.com/nqrm/go_bot/internal/repository/expenses/model"
)

var _ def.ExpenseRepository = (*repository)(nil)

type repository struct {
	data map[string]*repoModel.Expense
	m    sync.RWMutex
}

func NewRepository() *repository {
	return &repository{
		data: make(map[string]*repoModel.Expense),
	}
}

func (r *repository) Create(_ context.Context, uuid string, info *model.ExpenseInfo) error {
	r.m.Lock()
	defer r.m.Unlock()

	r.data[uuid] = &repoModel.Expense{
		UUID:      uuid,
		CreatedAt: info.CreatedAt,
		Category:  info.Category,
		Amount:    info.Amount,
	}

	return nil
}

func (r *repository) Get(_ context.Context, uuid string) (*model.Expense, error) {
	r.m.RLock()
	defer r.m.RUnlock()

	expenses, ok := r.data[uuid]
	if !ok {
		return nil, nil
	}

	return (*model.Expense)(expenses), nil
}
