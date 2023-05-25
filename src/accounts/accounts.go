package accounts

import (
	"time"
)

type Account struct {
	Id        int       `json:"id"`
	Document  string    `json:"document"`
	CreatedAt time.Time `json:"-"`
}
