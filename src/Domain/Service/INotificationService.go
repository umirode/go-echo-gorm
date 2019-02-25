package Service

import "github.com/umirode/go-rest/src/Domain/Model/Entity"

type INotificationService interface {
	SendToSingleUser(notification *Entity.Notification, user *Entity.User) error
	SendToAllUsers(notification *Entity.Notification) error
}
