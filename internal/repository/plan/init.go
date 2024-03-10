package plan

import (
	dto "github.com/Marvit-Solutions/csw-golang/internal/domain/response"

	"gorm.io/gorm"
)

type PlanRepository interface {
	ListPlan() ([]dto.PlanResponse, error)
	GetTop3Plan() ([]dto.SubPlanTop3Response, error)
}

type planRepository struct {
	db *gorm.DB
}

func NewPlanRepository(db *gorm.DB) PlanRepository {
	return &planRepository{db}
}
