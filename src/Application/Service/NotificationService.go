package Service

import (
	"github.com/NaySoftware/go-fcm"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Model/ValueObject"
	"github.com/umirode/go-rest/src/Domain/Repository"
)

type NotificationService struct {
	notificationTokenRepository Repository.INotificationTokenRepository
	birthdayRepository          Repository.IBirthdayRepository

	fcmClient *fcm.FcmClient
}

func NewNotificationService(fcmAPIKey string, notificationTokenRepository Repository.INotificationTokenRepository) *NotificationService {
	return &NotificationService{
		notificationTokenRepository: notificationTokenRepository,
		fcmClient:                   fcm.NewFcmClient(fcmAPIKey),
	}
}

func (s *NotificationService) SendToAllUsers(notification *ValueObject.Notification) error {
	tokens, err := s.notificationTokenRepository.FindAll()
	if err != nil {
		return err
	}

	if len(tokens) == 0 {
		return nil
	}

	s.fcmClient.SetNotificationPayload(&fcm.NotificationPayload{
		Title: notification.Title,
		Body:  notification.Message,
	})

	s.fcmClient.NewFcmRegIdsMsg(func() []string {
		newTokens := make([]string, 0)
		for _, token := range tokens {
			newTokens = append(newTokens, token.Token)
		}

		return newTokens
	}(), nil)

	status, err := s.fcmClient.Send()

	if err == nil {
		status.PrintResults()
	} else {
		return err
	}

	return nil
}

func (s *NotificationService) SendToSingleUser(notification *ValueObject.Notification, user *Entity.User) error {
	tokens, err := s.notificationTokenRepository.FindAllByUser(user)
	if err != nil {
		return err
	}

	if len(tokens) == 0 {
		return nil
	}

	s.fcmClient.SetNotificationPayload(&fcm.NotificationPayload{
		Title: notification.Title,
		Body:  notification.Message,
	})

	s.fcmClient.NewFcmRegIdsMsg(func() []string {
		newTokens := make([]string, 0)
		for _, token := range tokens {
			newTokens = append(newTokens, token.Token)
		}

		return newTokens
	}(), nil)

	status, err := s.fcmClient.Send()

	if err == nil {
		status.PrintResults()
	} else {
		return err
	}

	return nil
}
