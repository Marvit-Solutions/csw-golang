package plan

import (
	"csw-golang/internal/domain/entity/dto"
)

func (pc *planUsecase) ListPlan() ([]dto.PlanResponse, error) {
	ListPlan, err := pc.paketRepo.ListPlan()
	if err != nil {
		return ListPlan, err
	}

	return ListPlan, err
}

func (pc *planUsecase) GetTop3Plan() ([]dto.SubPlanTop3Response, error) {
	GetTop3Plan, err := pc.paketRepo.GetTop3Plan()
	if err != nil {
		return nil, err
	}

	return GetTop3Plan, err
}

// func (pc *planUsecase) CreatePaket(request dto.PaketRequest) (dto.PlanResponse, error) {
// 	createPaket, err := pc.paketRepo.CreatePaket(request)
// 	if err != nil {
// 		return createPaket, err
// 	}

// 	return createPaket, err
// }

// func (pc *planUsecase) UpdatePaket(request dto.PaketRequest, id string) (dto.PlanResponse, error) {
// 	updatePaket, err := pc.paketRepo.UpdatePaket(request, id)
// 	if err != nil {
// 		return updatePaket, err
// 	}
// 	return updatePaket, err
// }

// func (pc *planUsecase) DeletePaket(id string) (dto.PlanResponse, error) {
// 	deletePaket, err := pc.paketRepo.DeletePaket(id)
// 	if err != nil {
// 		return deletePaket, err
// 	}

// 	return deletePaket, err
// }

// func (pc *planUsecase) ListSubPlan(idPaket string) ([]dto.SubPlanResponse, error) {
// 	listSubPlan, err := pc.paketRepo.ListSubPlan(idPaket)
// 	if err != nil {
// 		return listSubPlan, err
// 	}

// 	return listSubPlan, err
// }

// func (pc *planUsecase) CreateSubPlan(request dto.SubPlanRequest) (dto.SubPlanResponse, error) {
// 	createSubPlan, err := pc.paketRepo.CreateSubPlan(request)
// 	if err != nil {
// 		return createSubPlan, err
// 	}

// 	return createSubPlan, err
// }

// func (pc *planUsecase) UpdateSubPlan(request dto.SubPlanRequest, id string) (dto.SubPlanResponse, error) {
// 	updateSubPlan, err := pc.paketRepo.UpdateSubPlan(request, id)
// 	if err != nil {
// 		return updateSubPlan, err
// 	}

// 	return updateSubPlan, err
// }

// func (pc *planUsecase) DeleteSubPlan(id string) (dto.SubPlanResponse, error) {
// 	deleteSubPlan, err := pc.paketRepo.DeleteSubPlan(id)
// 	if err != nil {
// 		return deleteSubPlan, err
// 	}

// 	return deleteSubPlan, err
// }

// func (pc *planUsecase) GetTopSubPlan() ([]dto.TopSubPlanResponse, error) {
// 	getTopSubPlan, err := pc.paketRepo.GetTopSubPlan()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return getTopSubPlan, err
// }
