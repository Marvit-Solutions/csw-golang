package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type PlanRepository interface {
	Create(model *model.Plan, tx *gorm.DB) (*model.Plan, error)
	Update(model *model.Plan, tx *gorm.DB) error
	Delete(model *model.Plan, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.Plan, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.Plan, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.Plan) error
}
