package auth

import (
	"csw-golang/internal/domain/entity/dto"
)

func (pc *paketUsecase) ListPaket() ([]dto.PaketResponse, error) {
	listPaket, err := pc.paketRepo.ListPaket()
	if err != nil {
		return listPaket, err
	}

	return listPaket, err
}

func (pc *paketUsecase) ListSubPaket(idPaket string) ([]dto.SubPaketResponse, error) {
	listSubPaket, err := pc.paketRepo.ListSubPaket(idPaket)
	if err != nil {
		return listSubPaket, err
	}

	return listSubPaket, err
}
