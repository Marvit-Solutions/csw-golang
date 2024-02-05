package pretest

import (
	"csw-golang/internal/domain/entity/datastruct"

	"gorm.io/gorm"
)

type PretestRepo interface {
	GetAllPretests() (error, []datastruct.Modules)
}

type pretestRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) PretestRepo {
	return &pretestRepo{
		db,
	}
}
