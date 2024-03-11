package plan

import (
	dto "github.com/Marvit-Solutions/csw-golang/internal/domain/response"
)

func (pc *planUsecase) ListPlan() ([]dto.PlanResponse, error) {
	data, err := pc.planRepo.ListPlan()
	if err != nil {
		return data, err
	}

	return data, err
}

func (pc *planUsecase) GetTop3Plan() ([]dto.SubPlanTop3Response, error) {
	data, err := pc.planRepo.GetTop3Plan()
	if err != nil {
		return nil, err
	}

	return data, err
}
