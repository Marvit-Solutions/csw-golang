package plan

import (
	"csw-golang/internal/domain/entity/dto"
	pr "csw-golang/internal/domain/repository/plan"
)

type PlanUsecase interface {
	ListPlan() ([]dto.PlanResponse, error)
	GetTop3Plan() ([]dto.SubPlanTop3Response, error)
	// CreatePaket(request dto.PaketRequest) (dto.PaketResponse, error)
	// UpdatePaket(request dto.PaketRequest, id string) (dto.PaketResponse, error)
	// DeletePaket(id string) (dto.PaketResponse, error)

	// ListSubPaket(idPaket string) ([]dto.SubPaketResponse, error)
	// CreateSubPaket(request dto.SubPaketRequest) (dto.SubPaketResponse, error)
	// UpdateSubPaket(request dto.SubPaketRequest, id string) (dto.SubPaketResponse, error)
	// DeleteSubPaket(id string) (dto.SubPaketResponse, error)
}

type planUsecase struct {
	paketRepo pr.PlanRepo
}

func New(paketRepo pr.PlanRepo) *planUsecase {
	return &planUsecase{
		paketRepo,
	}
}
