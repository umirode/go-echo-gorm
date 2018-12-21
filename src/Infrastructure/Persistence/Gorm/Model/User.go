package Model

type User struct {
	ID uint `gorm:"primary_key" json:"id"`

	Email        string `gorm:"not null;unique;size:255" json:"email"`
	PasswordHash string `gorm:"not null;size:255" json:"-"`

	Birthdays []*Birthday `gorm:"foreignkey:UserID"`
}
