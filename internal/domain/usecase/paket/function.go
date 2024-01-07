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

func (pc *paketUsecase) CreatePaket(request dto.PaketRequest) (dto.PaketResponse, error) {
	createPaket, err := pc.paketRepo.CreatePaket(request)
	if err != nil {
		return createPaket, err
	}

	return createPaket, err
}

func (pc *paketUsecase) UpdatePaket(request dto.PaketRequest, id string) (dto.PaketResponse, error) {
	updatePaket, err := pc.paketRepo.UpdatePaket(request, id)
	if err != nil {
		return updatePaket, err
	}
	return updatePaket, err
}

func (pc *paketUsecase) DeletePaket(id string) (dto.PaketResponse, error) {
	deletePaket, err := pc.paketRepo.DeletePaket(id)
	if err != nil {
		return deletePaket, err
	}

	return deletePaket, err
}

func (pc *paketUsecase) ListSubPaket(idPaket string) ([]dto.SubPaketResponse, error) {
	listSubPaket, err := pc.paketRepo.ListSubPaket(idPaket)
	if err != nil {
		return listSubPaket, err
	}

	return listSubPaket, err
}

func (pc *paketUsecase) CreateSubPaket(request dto.SubPaketRequest) (dto.SubPaketResponse, error) {
	createSubPaket, err := pc.paketRepo.CreateSubPaket(request)
	if err != nil {
		return createSubPaket, err
	}

	return createSubPaket, err
}

func (pc *paketUsecase) UpdateSubPaket(request dto.SubPaketRequest, id string) (dto.SubPaketResponse, error) {
	updateSubPaket, err := pc.paketRepo.UpdateSubPaket(request, id)
	if err != nil {
		return updateSubPaket, err
	}

	return updateSubPaket, err
}

func (pc *paketUsecase) DeleteSubPaket(id string) (dto.SubPaketResponse, error) {
	deleteSubPaket, err := pc.paketRepo.DeleteSubPaket(id)
	if err != nil {
		return deleteSubPaket, err
	}

	return deleteSubPaket, err
}

func (pc *paketUsecase) GetTopSubPaket() ([]dto.TopSubPaketResponse, error) {
	getTopSubPaket, err := pc.paketRepo.GetTopSubPaket()
	if err != nil {
		return nil, err
	}

	return getTopSubPaket, err
}
