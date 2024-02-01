package paket

import (
	pc "csw-golang/internal/domain/usecase/paket"
)

type PlanHandler struct {
	paketUsecase pc.PlanUsecase
}

func New(planUsecase pc.PlanUsecase) *PlanHandler {
	return &PlanHandler{
		planUsecase,
	}
}
