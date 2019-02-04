package Entity

type Birthday struct {
	ID uint `gorm:"primary_key"`

	Name   string `gorm:"size:20"`
	Month  uint   `gorm:"size:2"`
	Number uint   `gorm:"size:2"`

	OwnerID uint // User
}
