package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type RoleRepository interface {
	Create(model *model.Role, tx *gorm.DB) (*model.Role, error)
	Update(model *model.Role, tx *gorm.DB) error
	Delete(model *model.Role, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.Role, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.Role, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.Role) error
}
