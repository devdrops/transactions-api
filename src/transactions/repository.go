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

func (t *TransactionsRepository) Create(aId int, oId int, a float32) error {
	st, err := t.db.Conn.Prepare("INSERT INTO transactions(account_id, operation_id, amount) VALUES( $1, $2, $3)")
	if err != nil {
		return err
	}
	defer st.Close()

	if _, err := st.Exec(aId, oId, a); err != nil {
		return err
	}

	return nil
}
