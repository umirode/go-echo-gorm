package Service

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Service"
)

type BirthdayNotificationService struct {
	notificationService Service.INotificationService
	birthdayService     Service.IBirthdayService
}

func NewBirthdayNotificationService(notificationService Service.INotificationService, birthdayService Service.IBirthdayService) *BirthdayNotificationService {
	return &BirthdayNotificationService{
		notificationService: notificationService,
		birthdayService:     birthdayService,
	}
}

func (s *BirthdayNotificationService) SendNotificationsAboutBirthdays() error {
	todayBirthdays, err := s.birthdayService.GetAllWhichWillBeToday()
	if err != nil {
		return err
	}
	tomorrowBirthdays, err := s.birthdayService.GetAllWhichWillBeTomorrow()
	if err != nil {
		return err
	}

	for _, birthday := range todayBirthdays {
		notification := &Entity.Notification{
			Title:   "Birthday today",
			Message: "",
		}

		err := s.notificationService.SendToSingleUser(notification, &Entity.User{ID: birthday.ID})
		if err != nil {
			return err
		}
	}

	for _, birthday := range tomorrowBirthdays {
		notification := &Entity.Notification{
			Title:   "Birthday tomorrow",
			Message: "",
		}

		err := s.notificationService.SendToSingleUser(notification, &Entity.User{ID: birthday.ID})
		if err != nil {
			return err
		}
	}
	return nil
}
