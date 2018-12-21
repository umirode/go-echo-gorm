package Entity

type User struct {
	ID uint

	Email        string
	PasswordHash string

	Birthdays []*Birthday
}
