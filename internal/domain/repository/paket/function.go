package auth

import (
	"csw-golang/internal/domain/entity/dto"
)

func (pr *paketRepo) ListPaket() ([]dto.PaketResponse, error) {
	var listPaket []dto.PaketResponse

	err := pr.db.Raw("SELECT id, nama_paket, deskripsi_paket FROM pakets").Scan(&listPaket).Error

	if err != nil {
		return listPaket, err
	}

	return listPaket, nil
}

func (pr *paketRepo) ListSubPaket(idPaket string) ([]dto.SubPaketResponse, error) {
	var listSubPaket []dto.SubPaketResponse

	err := pr.db.Raw("SELECT s.id, p.nama_paket, s.nama_sub_paket, s.deskripsi_sub_paket, s.harga FROM sub_pakets s, pakets p WHERE s.id_paket = p.id AND id_paket = ?", idPaket).Scan(&listSubPaket).Error

	if err != nil {
		return listSubPaket, err
	}

	return listSubPaket, nil
}
