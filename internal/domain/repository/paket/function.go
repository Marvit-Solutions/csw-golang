package auth

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
)

func (pr *paketRepo) ListPaket() ([]dto.PaketResponse, error) {
	var paketList []datastruct.Paket

	err := pr.db.Preload("SubPaket").Preload("SubPaket.SubPaketDetail").Find(&paketList).Error
	if err != nil {
		return nil, err
	}

	var paketResponses []dto.PaketResponse
	for _, paket := range paketList {
		paketResponse := dto.PaketResponse{
			ID:        paket.ID,
			NamaPaket: paket.Nama,
		}

		for _, subPaket := range paket.SubPaket {
			subPaketResponse := dto.SubPaket{
				ID:           subPaket.ID,
				NamaSubPaket: subPaket.Nama,
				Harga:        subPaket.Harga,
				SubPaketDetail: dto.SubPaketDetail{
					ID:          subPaket.SubPaketDetail.ID,
					GrupPejuang: subPaket.SubPaketDetail.GrupPejuang,
					SoalLatihan: subPaket.SubPaketDetail.SoalLatihan,
					Akses:       subPaket.SubPaketDetail.Akses,
					Modul:       subPaket.SubPaketDetail.Modul,
					TryOut:      subPaket.SubPaketDetail.TryOut,
					Zoom:        subPaket.SubPaketDetail.Zoom,
				},
			}

			paketResponse.SubPaket = append(paketResponse.SubPaket, subPaketResponse)
		}

		paketResponses = append(paketResponses, paketResponse)
	}

	return paketResponses, nil
}

// func (pr *paketRepo) CreatePaket(request dto.PaketRequest) (dto.PaketResponse, error) {
// 	var paket dto.PaketResponse

// 	err := pr.db.Raw("INSERT INTO pakets (nama_paket, deskripsi_paket) VALUES (?, ?) RETURNING id,nama_paket,deskripsi_paket", request.NamaPaket, request.DeskripsiPaket).Scan(&paket).Error

// 	if err != nil {
// 		return paket, err
// 	}

// 	return paket, nil

// }

// func (pr *paketRepo) UpdatePaket(request dto.PaketRequest, id string) (dto.PaketResponse, error) {
// 	var paket dto.PaketResponse

// 	err := pr.db.Raw("UPDATE pakets SET nama_paket = ?, deskripsi_paket = ? WHERE id = ? RETURNING id,nama_paket,deskripsi_paket", request.NamaPaket, request.DeskripsiPaket, id).Scan(&paket).Error

// 	if err != nil {
// 		return paket, err
// 	}

// 	return paket, nil

// }

// func (pr *paketRepo) DeletePaket(id string) (dto.PaketResponse, error) {
// 	var paket dto.PaketResponse

// 	err := pr.db.Raw("DELETE FROM pakets WHERE id = ? RETURNING id,nama_paket,deskripsi_paket", id).Scan(&paket).Error
// 	if err != nil {
// 		return paket, err
// 	}

// 	return paket, nil
// }

// func (pr *paketRepo) ListSubPaket(idPaket string) ([]dto.SubPaketResponse, error) {
// 	var listSubPaket []dto.SubPaketResponse

// 	err := pr.db.Raw("SELECT s.id, p.nama_paket, s.nama_sub_paket, s.deskripsi_sub_paket, s.harga FROM sub_pakets s, pakets p WHERE s.id_paket = p.id AND id_paket = ?", idPaket).Scan(&listSubPaket).Error

// 	if err != nil {
// 		return listSubPaket, err
// 	}

// 	return listSubPaket, nil
// }

// func (pr *paketRepo) CreateSubPaket(request dto.SubPaketRequest) (dto.SubPaketResponse, error) {
// 	var subPaket dto.SubPaketResponse

// 	err := pr.db.Raw("INSERT INTO sub_pakets (id_paket, nama_sub_paket, deskripsi_sub_paket, harga) VALUES (?, ?, ?, ?) RETURNING id, nama_sub_paket, deskripsi_sub_paket, harga", request.IDPaket, request.NamaSubPaket, request.DeskripsiSubPaket, request.Harga).Scan(&subPaket).Error

// 	if err != nil {
// 		return subPaket, err
// 	}

// 	return subPaket, nil
// }

// func (pr *paketRepo) UpdateSubPaket(request dto.SubPaketRequest, id string) (dto.SubPaketResponse, error) {
// 	var subPaket dto.SubPaketResponse

// 	err := pr.db.Raw("UPDATE sub_pakets SET nama_sub_paket = ?, deskripsi_sub_paket = ?, harga = ? WHERE id = ? RETURNING id, nama_sub_paket, deskripsi_sub_paket, harga", request.NamaSubPaket, request.DeskripsiSubPaket, request.Harga, id).Scan(&subPaket).Error

// 	if err != nil {
// 		return subPaket, err
// 	}

// 	return subPaket, nil
// }

// func (pr *paketRepo) DeleteSubPaket(id string) (dto.SubPaketResponse, error) {
// 	var subPaket dto.SubPaketResponse

// 	err := pr.db.Raw("DELETE FROM sub_pakets WHERE id = ? RETURNING id, nama_sub_paket, deskripsi_sub_paket, harga", id).Scan(&subPaket).Error

// 	if err != nil {
// 		return subPaket, err
// 	}

// 	return subPaket, nil
// }

// func (pr *paketRepo) GetTopSubPaket() ([]dto.TopSubPaketResponse, error) {
// 	var ListTopSubPaket []dto.TopSubPaketResponse

// 	err := pr.db.Raw("SELECT s.sub_paket_id, sp.nama_sub_paket , COUNT(s.id) AS Total FROM subscriptions s, sub_pakets sp WHERE s.sub_paket_id = sp.id GROUP BY s.sub_paket_id, sp.nama_sub_paket ORDER BY Total DESC LIMIT 3").Scan(&ListTopSubPaket).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return ListTopSubPaket, nil
// }
