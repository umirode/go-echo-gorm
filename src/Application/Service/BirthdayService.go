package Service

import (
	"github.com/umirode/go-rest/src/Domain/Error"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Repository"
	"github.com/umirode/go-rest/src/Domain/Service/DTO"
)

type BirthdayService struct {
	birthdayRepository Repository.IBirthdayRepository
}

func NewBirthdayService(birthdayRepository Repository.IBirthdayRepository) *BirthdayService {
	return &BirthdayService{
		birthdayRepository: birthdayRepository,
	}
}

func (s *BirthdayService) GetAllForUser(user *Entity.User) ([]*Entity.Birthday, error) {
	users, err := s.birthdayRepository.FindAllByUser(user)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *BirthdayService) GetOneById(id uint) (*Entity.Birthday, error) {
	user, err := s.birthdayRepository.FindOneById(id)
	if err != nil {
		return nil, Error.NewNotFoundError()
	}

	return user, nil
}

func (s *BirthdayService) Create(birthdayDTO *DTO.BirthdayDTO, user *Entity.User) error {
	birthday := &Entity.Birthday{
		Name: birthdayDTO.Name,
		Date: birthdayDTO.Date,
		User: user,
	}

	err := s.birthdayRepository.Save(birthday)

	if err != nil {
		return err
	}

	return nil
}

func (s *BirthdayService) Update(birthday *Entity.Birthday, birthdayDTO *DTO.BirthdayDTO, user *Entity.User) error {
	if birthday.User.ID != user.ID {
		return Error.NewAccessError()
	}

	birthday.Name = birthdayDTO.Name
	birthday.Date = birthdayDTO.Date

	err := s.birthdayRepository.Save(birthday)

	if err != nil {
		return err
	}

	return nil
}

func (s *BirthdayService) Delete(birthday *Entity.Birthday, user *Entity.User) error {
	if birthday.User.ID != user.ID {
		return Error.NewAccessError()
	}

	err := s.birthdayRepository.Delete(birthday)

	if err != nil {
		return err
	}

	return nil
}
