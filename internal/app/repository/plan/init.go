package plan

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/response"

	"gorm.io/gorm"
)

type PlanRepository interface {
	ListPlan() ([]response.PlanResponse, error)
	GetTop3Plan() ([]response.SubPlanTop3Response, error)
}

type planRepository struct {
	db *gorm.DB
}

func NewPlanRepository(db *gorm.DB) PlanRepository {
	return &planRepository{db}
}
