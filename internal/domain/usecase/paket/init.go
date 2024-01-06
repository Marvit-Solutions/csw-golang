package auth

import (
	"csw-golang/internal/domain/entity/dto"
	pr "csw-golang/internal/domain/repository/paket"
)

type PaketUsecase interface {
	ListPaket() ([]dto.PaketResponse, error)
	CreatePaket(request dto.PaketRequest) (dto.PaketResponse, error)
	UpdatePaket(request dto.PaketRequest, id string) (dto.PaketResponse, error)
	DeletePaket(id string) (dto.PaketResponse, error)

	ListSubPaket(idPaket string) ([]dto.SubPaketResponse, error)
	CreateSubPaket(request dto.SubPaketRequest) (dto.SubPaketResponse, error)
	UpdateSubPaket(request dto.SubPaketRequest, id string) (dto.SubPaketResponse, error)
	DeleteSubPaket(id string) (dto.SubPaketResponse, error)
	GetTopSubPaket() ([]dto.TopSubPaketResponse, error)
}

type paketUsecase struct {
	paketRepo pr.PaketRepo
}

func New(paketRepo pr.PaketRepo) *paketUsecase {
	return &paketUsecase{
		paketRepo,
	}
}
