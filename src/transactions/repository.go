package transactions

import (
	"transactions-api/pkg/database"
)

type TransactionsRepository struct {
	db *database.Database
}

func NewRepository(dbC *database.Database) *TransactionsRepository {
	return &TransactionsRepository{
		db: dbC,
	}
}

func (a *TransactionsRepository) Create(doc string) error {
	st, err := a.db.Conn.Prepare("INSERT INTO transactions(document) VALUES( $1)")
	if err != nil {
		return err
	}
	defer st.Close()

	if _, err := st.Exec(doc); err != nil {
		return err
	}

	return nil
}
