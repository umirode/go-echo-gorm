package Entity

import (
	"time"
)

type Birthday struct {
	ID uint

	Name string
	Date time.Time

	OwnerID uint // User
}
