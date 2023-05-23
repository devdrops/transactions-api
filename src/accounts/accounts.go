package accounts

import (
	"time"
)

type Account struct {
	Id        int
	Document  string
	CreatedAt time.Time
}
