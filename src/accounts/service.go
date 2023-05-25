package accounts

import (
	"transactions-api/pkg/database"
)

type AccountService struct {
	r *AccountsRepository
}

func NewService() (*AccountService, error) {
	db, err := database.New()
	if err != nil {
		return nil, err
	}

	return &AccountService{
		r: NewRepository(db),
	}, nil
}

func (s *AccountService) Create(a Account) error {
	return s.r.Create(a.Document)
}

func (s *AccountService) Get(a Account) (Account, error) {
	return s.r.Get(a.Id)
}
