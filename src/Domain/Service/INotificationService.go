package Service

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Model/ValueObject"
)

type INotificationService interface {
	SendToSingleUser(notification *ValueObject.Notification, user *Entity.User) error
	SendToAllUsers(notification *ValueObject.Notification) error
}
