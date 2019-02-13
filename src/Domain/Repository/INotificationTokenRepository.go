package Repository

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

type INotificationTokenRepository interface {
	Save(token *Entity.NotificationToken) error
	Delete(token *Entity.NotificationToken) error

	FindAll() ([]*Entity.NotificationToken, error)
	FindAllByUser(user *Entity.User) ([]*Entity.NotificationToken, error)
	FindOneByTokenAndUser(token string, user *Entity.User) (*Entity.NotificationToken, error)
}
