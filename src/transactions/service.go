package transactions

import (
	"transactions-api/pkg/database"
)

type TransactionService struct {
	r *TransactionsRepository
}

func NewService() (*TransactionService, error) {
	db, err := database.New()
	if err != nil {
		return nil, err
	}

	return &TransactionService{
		r: NewRepository(db),
	}, nil
}

func (s *TransactionService) Create(t Transaction) error {
	return s.r.Create(t.AccountId, t.OperationId, t.Amount)
}
