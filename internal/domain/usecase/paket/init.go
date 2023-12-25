package auth

import (
	"csw-golang/internal/domain/entity/dto"
	pr "csw-golang/internal/domain/repository/paket"
)

type PaketUsecase interface {
	ListPaket() ([]dto.PaketResponse, error)
	ListSubPaket(idPaket string) ([]dto.SubPaketResponse, error)
}

type paketUsecase struct {
	paketRepo pr.PaketRepo
}

func New(paketRepo pr.PaketRepo) *paketUsecase {
	return &paketUsecase{
		paketRepo,
	}
}
