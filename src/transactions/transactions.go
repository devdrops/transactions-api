package transactions

import (
	"time"
)

type Transaction struct {
	Id          int
	AccountId   int
	OperationId int
	Amount      float32
	CreatedAt   time.Time
}
