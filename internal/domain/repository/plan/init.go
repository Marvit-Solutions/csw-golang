package plan

import (
	"csw-golang/internal/domain/entity/dto"

	"gorm.io/gorm"
)

type PlanRepo interface {
	ListPlan() ([]dto.PlanResponse, error)
	GetTop3Plan() ([]dto.SubPlanTop3Response, error)
	// CreatePaket(request dto.PaketRequest) (dto.PlanResponse, error)
	// UpdatePaket(request dto.PaketRequest, id string) (dto.PlanResponse, error)
	// DeletePaket(id string) (dto.PlanResponse, error)
	// Sub Paket
	// CreateSubPlan(request dto.SubPlanRequest) (dto.SubPlanResponse, error)
	// UpdateSubPlan(request dto.SubPlanRequest, id string) (dto.SubPlanResponse, error)
	// DeleteSubPlan(id string) (dto.SubPlanResponse, error)
}

type planRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) PlanRepo {
	return &planRepo{
		db,
	}
}
