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

func (s *BirthdayService) GetOneByIdAndUser(id uint, user *Entity.User) (*Entity.Birthday, error) {
	birthday, err := s.birthdayRepository.FindOneByIdAndUser(id, user)
	if err != nil {
		return nil, err
	}

	if birthday == nil {
		return nil, Error.NewNotFoundError()
	}

	return birthday, nil
}

func (s *BirthdayService) GetAllByUser(user *Entity.User) ([]*Entity.Birthday, error) {
	birthdays, err := s.birthdayRepository.FindAllByUser(user)
	if err != nil {
		return nil, err
	}

	return birthdays, nil
}

func (s *BirthdayService) Create(birthdayDTO *DTO.BirthdayDTO, user *Entity.User) error {
	birthday := &Entity.Birthday{
		Name:    birthdayDTO.Name,
		Date:    birthdayDTO.Date,
		OwnerID: user.ID,
	}

	err := s.birthdayRepository.Save(birthday)

	if err != nil {
		return err
	}

	return nil
}

func (s *BirthdayService) Update(birthday *Entity.Birthday, birthdayDTO *DTO.BirthdayDTO, user *Entity.User) error {
	if birthday.OwnerID != user.ID {
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
	if birthday.OwnerID != user.ID {
		return Error.NewAccessError()
	}

	err := s.birthdayRepository.Delete(birthday)

	if err != nil {
		return err
	}

	return nil
}
