package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type ClassPlanUserRepository interface {
	Create(model *model.ClassPlanUser, tx *gorm.DB) (*model.ClassPlanUser, error)
	Update(model *model.ClassPlanUser, tx *gorm.DB) error
	Delete(model *model.ClassPlanUser, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.ClassPlanUser, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.ClassPlanUser, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.ClassPlanUser) error
}
