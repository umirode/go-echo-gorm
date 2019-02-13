package Entity

type Birthday struct {
	ID uint `gorm:"primary_key"`

	Name   string `gorm:"size:20"`
	Month  uint   `gorm:"size:2"`
	Day uint   `gorm:"size:2"`
	Year   uint   `gorm:"size:4"`

	OwnerID uint // User
}
