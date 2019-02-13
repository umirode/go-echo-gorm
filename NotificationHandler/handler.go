package NotificationHandler

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/src/Application/Service"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Infrastructure/Persistence/Gorm/Repository"
)

type NotificationHandler struct {
	database  *gorm.DB
	fcmAPIKey string
}

func NewNotificationHandler(database *gorm.DB, fcmAPIKey string) *NotificationHandler {
	return &NotificationHandler{
		database:  database,
		fcmAPIKey: fcmAPIKey,
	}
}

func (n *NotificationHandler) Run() {
	notificationService := Service.NewNotificationService(n.fcmAPIKey, Repository.NewNotificationTokenRepository(n.database))

	notificationService.SendToAllUsers(&Entity.Notification{
		Title:   "Ты пидор",
		Message: "пидор пидор пидор пидор пидор пидор пидор ",
	})
}
