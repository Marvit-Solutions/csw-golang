package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type ClassPlanTypeRepository interface {
	Create(model *model.ClassPlanType, tx *gorm.DB) (*model.ClassPlanType, error)
	Update(model *model.ClassPlanType, tx *gorm.DB) error
	Delete(model *model.ClassPlanType, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.ClassPlanType, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.ClassPlanType, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.ClassPlanType) error
}
