package plan

import (
	"csw-golang/internal/domain/entity/dto"
	pr "csw-golang/internal/domain/repository/plan"
)

type PlanUsecase interface {
	ListPlan() ([]dto.PlanResponse, error)
	GetTop3Plan() ([]dto.SubPlanTop3Response, error)
}

type planUsecase struct {
	planRepo pr.PlanRepo
}

func New(planRepo pr.PlanRepo) *planUsecase {
	return &planUsecase{
		planRepo,
	}
}
