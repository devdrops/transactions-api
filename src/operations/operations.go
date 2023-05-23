package operations

import (
	"time"
)

type Operation struct {
	Id          int
	Description string
	CreatedAt   time.Time
}
