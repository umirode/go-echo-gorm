package Repository

import (
	"github.com/umirode/go-rest/src/Domain/Model/Entity"
)

type NotificationTokenRepository struct {
	*BaseRepository
}

func NewNotificationTokenRepository() *NotificationTokenRepository {
	return &NotificationTokenRepository{}
}

func (r *NotificationTokenRepository) FindOneByTokenAndUser(token string, user *Entity.User) (*Entity.NotificationToken, error) {
	notificationToken := &Entity.NotificationToken{}

	r.GetGormDB().Where("token = ? and owner_id = ?", token, user.ID).First(notificationToken)
	if r.GetGormDB().NewRecord(notificationToken) {
		return nil, nil
	}

	return notificationToken, nil
}

func (r *NotificationTokenRepository) FindAll() ([]*Entity.NotificationToken, error) {
	tokens := new([]*Entity.NotificationToken)

	r.GetGormDB().Find(tokens)

	return *tokens, nil
}

func (r *NotificationTokenRepository) FindAllByUser(user *Entity.User) ([]*Entity.NotificationToken, error) {
	tokens := new([]*Entity.NotificationToken)

	r.GetGormDB().Where("owner_id = ?", user.ID).Find(tokens)

	return *tokens, nil
}

func (r *NotificationTokenRepository) Save(token *Entity.NotificationToken) error {
	r.GetGormDB().Save(token)

	return nil
}

func (r *NotificationTokenRepository) Delete(token *Entity.NotificationToken) error {
	r.GetGormDB().Delete(token)

	return nil
}
