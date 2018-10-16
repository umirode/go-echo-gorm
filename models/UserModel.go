package models

type UserModel struct {
	IModel `json:"-"`

	BaseModel

	Name string `gorm:"size:255" json:"name"`
}

func (m *UserModel) TableName() string {
	return "users"
}
