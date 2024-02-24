package pretest

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"
	pr "csw-golang/internal/domain/repository/pretest"
)

type PretestUsecase interface {
	GetAllPretests() (error, []dto.GetAllPretestResponse)
	GetPretestById(pretestId string) (error, dto.Pretest)
	GetPretestReview(pretestId, status string) (error, dto.PretestReviewResponse)
	SubmitPretest(id string, req request.PretestSubmitRequest) error
	GradingPretest(id string, req request.PretestSubmitRequest) error
}

type pretestUsecase struct {
	pretestRepo pr.PretestRepo
}

func New(pretestRepo pr.PretestRepo) *pretestUsecase {
	return &pretestUsecase{
		pretestRepo,
	}
}
