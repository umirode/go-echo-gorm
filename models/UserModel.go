package models

type UserModel struct {
	BaseModel

	Name string `gorm:"size:255" json:"name"`
}

func (m *UserModel) TableName() string {
	return "users"
}
