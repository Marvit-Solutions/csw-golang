package plan

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/response"
)

func (pc *planUsecase) ListPlan() ([]response.PlanResponse, error) {
	data, err := pc.planRepo.ListPlan()
	if err != nil {
		return data, err
	}

	return data, err
}

func (pc *planUsecase) GetTop3Plan() ([]response.SubPlanTop3Response, error) {
	data, err := pc.planRepo.GetTop3Plan()
	if err != nil {
		return nil, err
	}

	return data, err
}
