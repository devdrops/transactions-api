package accounts

import (
	"transactions-api/pkg/database"
)

type AccountsRepository struct {
	db *database.Database
}

func NewRepository(dbC *database.Database) *AccountsRepository {
	return &AccountsRepository{
		db: dbC,
	}
}

func (a *AccountsRepository) Create(doc string) error {
	st, err := a.db.Conn.Prepare("INSERT INTO accounts(document) VALUES( $1)")
	if err != nil {
		return err
	}
	defer st.Close()

	if _, err := st.Exec(doc); err != nil {
		return err
	}

	return nil
}

func (a *AccountsRepository) Get(id int) (Account, error) {
	var acc Account
	err := a.db.Conn.QueryRow("SELECT id, document FROM accounts WHERE id = $1", id).
		Scan(&acc.Id, &acc.Document)
	if err != nil {
		return acc, err
	}

	return acc, nil
}
