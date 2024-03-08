package module

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"
	module "csw-golang/internal/domain/repository/modul"
)

type ModuleUsecase interface {
	GetListModules() ([]dto.ModuleResponse, error)
	GetSubjectsBySubmoduleID(submoduleID string) ([]dto.SubjectResponse, error)
	GetQuestionsByTestTypeID(testTypeID string) (dto.ExerciseResponse, error)
	GetTestReview(moduleTestID string) (dto.ReviewResultResponse, error)
	PostSubmittedTest(testTypeID, userID string, submittedQuiz request.UserSubmittedQuizRequest) error
	GetTop3EverySubject(userID string, subjectTypeID string) ([]dto.HistoryTop3ScoreResponse, error)
}

type moduleUsecase struct {
	moduleRepo module.ModuleRepo
}

func New(moduleRepo module.ModuleRepo) ModuleUsecase {
	return &moduleUsecase{
		moduleRepo,
	}
}
