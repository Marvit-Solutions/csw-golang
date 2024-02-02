package plan

import (
	"csw-golang/internal/domain/entity/dto"

	"gorm.io/gorm"
)

type PlanRepo interface {
	ListPlan() ([]dto.PlanResponse, error)
	GetTop3Plan() ([]dto.SubPlanTop3Response, error)
}

type planRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) PlanRepo {
	return &planRepo{
		db,
	}
}
