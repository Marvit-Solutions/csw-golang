package module

import (
	"csw-golang/internal/domain/entity/dto"

	"gorm.io/gorm"
)

type ModuleRepo interface {
	ListModule() ([]dto.ModuleResponse, error)
}

type moduleRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) ModuleRepo {
	return &moduleRepo{
		db,
	}
}
