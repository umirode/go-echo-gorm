package Notification

import (
	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
	"github.com/umirode/go-rest/src/Application/Service"
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
	birthdayService := Service.NewBirthdayService(Repository.NewBirthdayRepository())
	birthdayNotificationService := Service.NewBirthdayNotificationService(notificationService, birthdayService)

	gocron.Every(1).Day().At("8:00").Do(func() {
		err := birthdayNotificationService.SendNotificationsAboutBirthdays()
		if err != nil {
			logrus.Error(err)
		}
	})

	<-gocron.Start()
}
