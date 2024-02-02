package plan

import (
	pc "csw-golang/internal/domain/usecase/plan"
)

type PlanHandler struct {
	planUsecase pc.PlanUsecase
}

func New(planUsecase pc.PlanUsecase) *PlanHandler {
	return &PlanHandler{
		planUsecase,
	}
}
