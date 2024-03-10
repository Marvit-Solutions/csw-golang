package plan

import (
	dto "csw-golang/internal/domain/response"
)

func (pc *planUsecase) ListPlan() ([]dto.PlanResponse, error) {
	ListPlan, err := pc.planRepo.ListPlan()
	if err != nil {
		return ListPlan, err
	}

	return ListPlan, err
}

func (pc *planUsecase) GetTop3Plan() ([]dto.SubPlanTop3Response, error) {
	GetTop3Plan, err := pc.planRepo.GetTop3Plan()
	if err != nil {
		return nil, err
	}

	return GetTop3Plan, err
}
