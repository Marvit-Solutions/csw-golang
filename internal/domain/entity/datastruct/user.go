package datastruct

import "gorm.io/gorm"

type User struct {
	*gorm.Model `json:"-"`
	RoleId      uint
	Email       string     `json:"Email" form:"Email" validate:"required,email"`
	GoogleId    string     `json:"GoogleId" form:"GoogleId"`
	FacebookId  string     `json:"FacebookId" form:"FacebookId"`
	Username    string     `json:"Username" form:"Username" validate:"required"`
	Password    string     `json:"Password" form:"Password" validate:"required,min=8"`
	UserDetail  UserDetail `gorm:"foreignKey:UserId"`
}
