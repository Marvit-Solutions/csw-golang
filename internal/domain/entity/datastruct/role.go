package datastruct

import "gorm.io/gorm"

type Role struct {
	*gorm.Model

	Role  string `json:"Role" form:"Role"`
	Users []User
}
