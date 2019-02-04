package Entity

type User struct {
	ID uint `gorm:"primary_key"`

	Email        string `gorm:"size:500"`
	PasswordHash string `gorm:"size:40"`
}
