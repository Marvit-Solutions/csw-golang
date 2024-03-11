package plan

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/repository/plan"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/response"

	"gorm.io/gorm"
)

type PlanUsecase interface {
	ListPlan() ([]response.PlanResponse, error)
	GetTop3Plan() ([]response.SubPlanTop3Response, error)
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
