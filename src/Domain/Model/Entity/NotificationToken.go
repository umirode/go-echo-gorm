package Entity

type NotificationToken struct {
	ID uint `gorm:"primary_key"`

	Token string `gorm:"size:255"`

	OwnerID uint // User
}
