package Service

import (
	"github.com/umirode/go-rest/src/Domain/Error"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Repository"
	"github.com/umirode/go-rest/src/Domain/Service/DTO"
)

type NotificationTokenService struct {
	notificationTokenRepository Repository.INotificationTokenRepository
}

func NewNotificationTokenService(notificationTokenRepository Repository.INotificationTokenRepository) *NotificationTokenService {
	return &NotificationTokenService{
		notificationTokenRepository: notificationTokenRepository,
	}
}

func (s *NotificationTokenService) Create(tokenDTO *DTO.NotificationTokenDTO, user *Entity.User) error {
	token, err := s.notificationTokenRepository.FindOneByTokenAndUser(tokenDTO.Token, user)
	if err != nil {
		return err
	}
	if token != nil {
		return Error.NewAlreadyExistsError()
	}

	token = &Entity.NotificationToken{
		Token:   tokenDTO.Token,
		OwnerID: user.ID,
	}

	err = s.notificationTokenRepository.Save(token)

	if err != nil {
		return err
	}

	return nil
}

func (s *NotificationTokenService) Delete(token *Entity.NotificationToken, user *Entity.User) error {
	if token.OwnerID != user.ID {
		return Error.NewAccessError()
	}

	err := s.notificationTokenRepository.Delete(token)

	if err != nil {
		return err
	}

	return nil
}

func (s *NotificationTokenService) GetAllByUser(user *Entity.User) ([]*Entity.NotificationToken, error) {
	tokens, err := s.notificationTokenRepository.FindAllByUser(user)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
