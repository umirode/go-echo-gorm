package Entity

import (
	"time"
)

type Birthday struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time

	Name string
	Date time.Time

	User *User
}
