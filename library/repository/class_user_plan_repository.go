package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type ClassUserPlanRepository interface {
	Create(model *model.ClassUserPlan, tx *gorm.DB) (*model.ClassUserPlan, error)
	Update(model *model.ClassUserPlan, tx *gorm.DB) error
	Delete(model *model.ClassUserPlan, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.ClassUserPlan, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.ClassUserPlan, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.ClassUserPlan) error
}
