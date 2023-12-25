package auth

import (
	"csw-golang/internal/domain/entity/dto"

	"gorm.io/gorm"
)

type PaketRepo interface {
	ListPaket() ([]dto.PaketResponse, error)
	ListSubPaket(idPaket string) ([]dto.SubPaketResponse, error)
}

type paketRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) PaketRepo {
	return &paketRepo{
		db,
	}
}
