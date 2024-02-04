package pretest

import (
	"csw-golang/internal/domain/entity/dto"

	"gorm.io/gorm"
)

type PretestRepo interface {
	GetAllPretests() (error, dto.GetAllPretestResponse)
}

type pretestRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) PretestRepo {
	return &pretestRepo{
		db,
	}
}
