package auth

import (
	"csw-golang/internal/domain/entity/dto"

	"gorm.io/gorm"
)

type PaketRepo interface {
	// PAKET
	ListPaket() ([]dto.PaketResponse, error)
	CreatePaket(request dto.PaketRequest) (dto.PaketResponse, error)
	UpdatePaket(request dto.PaketRequest, id string) (dto.PaketResponse, error)
	DeletePaket(id string) (dto.PaketResponse, error)
	// SUB PAKET
	ListSubPaket(idPaket string) ([]dto.SubPaketResponse, error)
	CreateSubPaket(request dto.SubPaketRequest) (dto.SubPaketResponse, error)
	UpdateSubPaket(request dto.SubPaketRequest, id string) (dto.SubPaketResponse, error)
	DeleteSubPaket(id string) (dto.SubPaketResponse, error)
}

type paketRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) PaketRepo {
	return &paketRepo{
		db,
	}
}
