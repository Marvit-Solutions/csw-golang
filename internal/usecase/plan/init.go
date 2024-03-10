package plan

import (
	dto "csw-golang/internal/domain/response"
	"csw-golang/internal/repository/plan"

	"gorm.io/gorm"
)

type PlanUsecase interface {
	ListPlan() ([]dto.PlanResponse, error)
	GetTop3Plan() ([]dto.SubPlanTop3Response, error)
}

type planUsecase struct {
	planRepo plan.PlanRepository
}

func NewPlanUsecase(
	db *gorm.DB,
) PlanUsecase {
	return &planUsecase{
		planRepo: plan.NewPlanRepository(db),
	}
}
