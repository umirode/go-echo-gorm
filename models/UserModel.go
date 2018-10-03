package models

type UserModel struct {
	BaseModel

	Field string `gorm:"size:255" json:"field"`
}

func (m *UserModel) TableName() string {
	return "users"
}
