package repository

import (
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"

	"gorm.io/gorm"
)

//Defines the interface for model repository operations.
type ScheduleRepository interface {
	Create(model *model.Schedule, tx *gorm.DB) (*model.Schedule, error)
	Update(model *model.Schedule, tx *gorm.DB) error
	Delete(model *model.Schedule, tx *gorm.DB) error
	FindOneBy(criteria map[string]interface{}) (*model.Schedule, error)
	FindBy(criteria map[string]interface{}, page, size int) ([]*model.Schedule, error)
	Count(criteria map[string]interface{}) int
	CreateOrUpdateIndex(model *model.Schedule) error
}
