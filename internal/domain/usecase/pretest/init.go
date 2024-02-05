package pretest

import (
	"csw-golang/internal/domain/entity/dto"
	pr "csw-golang/internal/domain/repository/pretest"
)

type PretestUsecase interface {
	GetAllPretests() (error, []dto.GetAllPretestResponse)
}

type pretestUsecase struct {
	pretestRepo pr.PretestRepo
}

func New(pretestRepo pr.PretestRepo) *pretestUsecase {
	return &pretestUsecase{
		pretestRepo,
	}
}
