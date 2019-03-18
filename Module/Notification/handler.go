package Notification

import (
	"github.com/sirupsen/logrus"
	"github.com/umirode/go-rest/src/Application/Service"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Infrastructure/Persistence/Gorm/Repository"
)

type Handler struct {
	fcmAPIKey string
}

func NewHandler(fcmAPIKey string) *Handler {
	return &Handler{
		fcmAPIKey: fcmAPIKey,
	}
}

func (n *Handler) Run() {
	notificationService := Service.NewNotificationService(n.fcmAPIKey, Repository.NewNotificationTokenRepository())

	err := notificationService.SendToAllUsers(&Entity.Notification{
		Title:   "Test",
		Message: "Test Test",
	})
	if err != nil {
		logrus.Error(err)
	}
}
