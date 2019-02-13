package Repository

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

type BirthdayRepository struct {
	db *gorm.DB
}

func NewBirthdayRepository(db *gorm.DB) *BirthdayRepository {
	return &BirthdayRepository{
		db: db,
	}
}

func (r *BirthdayRepository) Save(birthday *Entity.Birthday) error {
	r.db.Save(birthday)

	return nil
}

func (r *BirthdayRepository) Delete(birthday *Entity.Birthday) error {
	r.db.Delete(birthday)

	return nil
}

func (r *BirthdayRepository) FindAllByUser(user *Entity.User) ([]*Entity.Birthday, error) {
	birthdays := new([]*Entity.Birthday)

	r.db.Where("owner_id = ?", user.ID).Find(birthdays)

	return *birthdays, nil
}

func (r *BirthdayRepository) FindOneByIdAndUser(id uint, user *Entity.User) (*Entity.Birthday, error) {
	birthday := &Entity.Birthday{}

	r.db.Where("id = ? and owner_id = ?", id, user.ID).First(birthday)
	if r.db.NewRecord(birthday) {
		return nil, nil
	}

	return birthday, nil
}
