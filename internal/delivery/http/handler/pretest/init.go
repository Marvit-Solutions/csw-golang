package pretest

import (
	pc "csw-golang/internal/domain/usecase/pretest"
)

type PretestHandler struct {
	pretestUsecase pc.PretestUsecase
}

func New(pretestUsecase pc.PretestUsecase) *PretestHandler {
	return &PretestHandler{
		pretestUsecase,
	}
}
