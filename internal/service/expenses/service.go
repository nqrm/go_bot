package expenses

import (
	"context"
	"log"

	"github.com/google/uuid"
	model "github.com/nqrm/go_bot/internal/model/expenses"
	"github.com/nqrm/go_bot/internal/repository"
	def "github.com/nqrm/go_bot/internal/service"
)

type service struct {
	expsensesRepository repository.ExpenseRepository
}

var _ def.ExpenseService = (*service)(nil)

func (s *service) Create(ctx context.Context, info *model.ExpenseInfo) (string, error) {
	expsensesUUID, err := uuid.NewUUID()
	if err != nil {
		log.Printf("ошибка генерации UUID: %v\n", err)
		return "", err
	}
	err = s.expsensesRepository.Create(ctx, expsensesUUID.String(), info)
	if err != nil {
		log.Printf("ошибка создания расхода: %v\n", err)
		return "", err
	}

	return expsensesUUID.String(), nil
}

func (s *service) Get(ctx context.Context, uuid string) (*model.Expense, error) {
	expense, err := s.expsensesRepository.Get(ctx, uuid)
	if err != nil {
		log.Printf("ошибка получения расхода: %v\n", err)
		return nil, nil
	}

	if expense == nil {
		log.Printf("расход с uuid %s не найден\n", err)
		return nil, nil
	}

	return expense, nil
}

func NewService(expsensesRepository repository.ExpenseRepository) *service {
	return &service{expsensesRepository: expsensesRepository}
}
