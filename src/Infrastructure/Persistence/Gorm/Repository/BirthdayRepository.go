package Repository

import (
	"time"

	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

type BirthdayRepository struct {
	BaseRepository
}

func NewBirthdayRepository() *BirthdayRepository {
	return &BirthdayRepository{}
}

func (r *BirthdayRepository) Save(birthday *Entity.Birthday) error {
	r.GetGormDB().Save(birthday)

	return nil
}

func (r *BirthdayRepository) Delete(birthday *Entity.Birthday) error {
	r.GetGormDB().Delete(birthday)

	return nil
}

func (r *BirthdayRepository) CountByUser(user *Entity.User) (uint, error) {
	count := new(uint)

	r.GetGormDB().Model(&Entity.Birthday{}).Where("owner_id = ?", user.ID).Count(count)

	return *count, nil
}

func (r *BirthdayRepository) FindAllByUser(user *Entity.User) ([]*Entity.Birthday, error) {
	birthdays := new([]*Entity.Birthday)

	r.GetGormDB().Where("owner_id = ?", user.ID).Find(birthdays)

	return *birthdays, nil
}

func (r *BirthdayRepository) FindOneByIdAndUser(id uint, user *Entity.User) (*Entity.Birthday, error) {
	birthday := &Entity.Birthday{}

	r.GetGormDB().Where("id = ? and owner_id = ?", id, user.ID).First(birthday)
	if r.GetGormDB().NewRecord(birthday) {
		return nil, nil
	}

	return birthday, nil
}

func (r *BirthdayRepository) FindAllWhichWillBeTomorrow() ([]*Entity.Birthday, error) {
	date := time.Now().AddDate(0, 0, 1)

	birthdays := new([]*Entity.Birthday)

	r.GetGormDB().Where("month = ? and day = ?", int(date.Month()), date.Day()).Find(birthdays)

	return *birthdays, nil
}

func (r *BirthdayRepository) FindAllWhichWillBeToday() ([]*Entity.Birthday, error) {
	date := time.Now()

	birthdays := new([]*Entity.Birthday)

	r.GetGormDB().Where("month = ? and day = ?", int(date.Month()), date.Day()).Find(birthdays)

	return *birthdays, nil
}
