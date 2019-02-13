package Repository

import (
	"github.com/jinzhu/gorm"
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

type NotificationTokenRepository struct {
	db *gorm.DB
}

func NewNotificationTokenRepository(db *gorm.DB) *NotificationTokenRepository {
	return &NotificationTokenRepository{
		db: db,
	}
}

func (r *NotificationTokenRepository) FindOneByTokenAndUser(token string, user *Entity.User) (*Entity.NotificationToken, error) {
	notificationToken := &Entity.NotificationToken{}

	r.db.Where("token = ? and owner_id = ?", token, user.ID).First(notificationToken)
	if r.db.NewRecord(notificationToken) {
		return nil, nil
	}

	return notificationToken, nil
}

func (r *NotificationTokenRepository) FindAll() ([]*Entity.NotificationToken, error) {
	tokens := new([]*Entity.NotificationToken)

	r.db.Find(tokens)

	return *tokens, nil
}

func (r *NotificationTokenRepository) FindAllByUser(user *Entity.User) ([]*Entity.NotificationToken, error) {
	tokens := new([]*Entity.NotificationToken)

	r.db.Where("owner_id = ?", user.ID).Find(tokens)

	return *tokens, nil
}

func (r *NotificationTokenRepository) Save(token *Entity.NotificationToken) error {
	r.db.Save(token)

	return nil
}

func (r *NotificationTokenRepository) Delete(token *Entity.NotificationToken) error {
	r.db.Delete(token)

	return nil
}
