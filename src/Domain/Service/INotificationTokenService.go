package Service

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Service/DTO"
)

type INotificationTokenService interface {
	GetAllByUser(user *Entity.User) ([]*Entity.NotificationToken, error)

	Create(tokenDTO *DTO.NotificationTokenDTO, user *Entity.User) error
	Delete(token *Entity.NotificationToken, user *Entity.User) error
}
