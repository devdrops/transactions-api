package transactions

import (
	"time"
)

type Transaction struct {
	Id          int       `json:"id"`
	AccountId   int       `json:"account_id"`
	OperationId int       `json:"operation_id"`
	Amount      float32   `json:"amount"`
	CreatedAt   time.Time `json:"-"`
}
