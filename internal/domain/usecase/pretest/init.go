package pretest

import (
	"csw-golang/internal/domain/entity/dto"
	pr "csw-golang/internal/domain/repository/pretest"
)

type PretestUsecase interface {
	GetAllPretests() (error, []dto.GetAllPretestResponse)
	GetPretestById(pretestId string) (error, dto.Pretest)
	GetPretestReview(pretestId, status string) (error, dto.Pretest)
}

type pretestUsecase struct {
	pretestRepo pr.PretestRepo
}

func New(pretestRepo pr.PretestRepo) *pretestUsecase {
	return &pretestUsecase{
		pretestRepo,
	}
}
