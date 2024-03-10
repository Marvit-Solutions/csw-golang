package plan

import (
	dto "github.com/Marvit-Solutions/csw-golang/internal/domain/response"
	"github.com/Marvit-Solutions/csw-golang/internal/repository/plan"

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
