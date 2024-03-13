package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type UserDetailRepository interface {
	Create(model *model.UserDetail, tx *gorm.DB) (*model.UserDetail, error)
	Update(model *model.UserDetail, tx *gorm.DB) error
	Delete(model *model.UserDetail, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.UserDetail, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.UserDetail, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.UserDetail) error
}
