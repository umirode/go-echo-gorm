package Service

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
	"github.com/umirode/go-rest/src/Domain/Repository"
	"github.com/umirode/go-rest/src/Domain/Service/DTO"
)

type BirthdayService struct {
	birthdayRepository Repository.IBirthdayRepository
}

func (s *BirthdayService) GetAllForUser(user *Entity.User) ([]*Entity.Birthday, error) {
	users, err := s.birthdayRepository.FindAllByUser(user)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *BirthdayService) Create(birthdayDTO DTO.BirthdayDTO, user *Entity.User) error {
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

func (s *BirthdayService) Update(birthday *Entity.Birthday, birthdayDTO DTO.BirthdayDTO) error {
	birthday.Name = birthdayDTO.Name
	birthday.Date = birthdayDTO.Date

	err := s.birthdayRepository.Save(birthday)

	if err != nil {
		return err
	}

	return nil
}

func (s *BirthdayService) Delete(birthday *Entity.Birthday) error {
	err := s.birthdayRepository.Delete(birthday)

	if err != nil {
		return err
	}

	return nil
}
