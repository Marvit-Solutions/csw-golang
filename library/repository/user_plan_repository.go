package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type UserPlanRepository interface {
	Create(model *model.UserPlan, tx *gorm.DB) (*model.UserPlan, error)
	Update(model *model.UserPlan, tx *gorm.DB) error
	Delete(model *model.UserPlan, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.UserPlan, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.UserPlan, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.UserPlan) error
}
