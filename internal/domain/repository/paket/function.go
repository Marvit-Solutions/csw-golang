package auth

import (
	"csw-golang/internal/domain/entity/dto"
)

func (pr *paketRepo) ListPaket() ([]dto.PaketResponse, error) {
	var listPaket []dto.PaketResponse

	err := pr.db.Find(&listPaket).Error

	if err != nil {
		return listPaket, err
	}

	return listPaket, nil
}

func (pr *paketRepo) ListSubPaket(idPaket string) ([]dto.SubPaketResponse, error) {
	var listSubPaket []dto.SubPaketResponse

	err := pr.db.Where("id_paket = ?", idPaket).Find(&listSubPaket).Error

	if err != nil {
		return listSubPaket, err
	}

	return listSubPaket, nil
}
