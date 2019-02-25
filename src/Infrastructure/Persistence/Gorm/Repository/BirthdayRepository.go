package Repository

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

type BirthdayRepository struct {
	*BaseRepository
}

func NewBirthdayRepository() *BirthdayRepository {
	return &BirthdayRepository{}
}

func (r *BirthdayRepository) Save(birthday *Entity.Birthday) error {
	r.GetDB().Save(birthday)

	return nil
}

func (r *BirthdayRepository) Delete(birthday *Entity.Birthday) error {
	r.GetDB().Delete(birthday)

	return nil
}

func (r *BirthdayRepository) CountByUser(user *Entity.User) (uint, error) {
	count := new(uint)

	r.GetDB().Where("owner_id = ?", user.ID).Count(count)

	return *count, nil
}

func (r *BirthdayRepository) FindAllByUser(user *Entity.User) ([]*Entity.Birthday, error) {
	birthdays := new([]*Entity.Birthday)

	r.GetDB().Where("owner_id = ?", user.ID).Find(birthdays)

	return *birthdays, nil
}

func (r *BirthdayRepository) FindOneByIdAndUser(id uint, user *Entity.User) (*Entity.Birthday, error) {
	birthday := &Entity.Birthday{}

	r.GetDB().Where("id = ? and owner_id = ?", id, user.ID).First(birthday)
	if r.GetDB().NewRecord(birthday) {
		return nil, nil
	}

	return birthday, nil
}
