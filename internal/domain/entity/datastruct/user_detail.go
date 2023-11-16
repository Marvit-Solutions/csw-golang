package datastruct

import "gorm.io/gorm"

type UserDetail struct {
	*gorm.Model `json:"-"`
	Name        string `json:"Name" form:"Name" validate:"required"`
	Phone       string `json:"Phone" form:"Phone" validate:"required,min=10,max=13,numeric"`
	UserId      uint   `json:"UserId" form:"UserId"`
}
