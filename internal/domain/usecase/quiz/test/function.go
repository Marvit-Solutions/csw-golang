package tests

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"
)

func (tc *testUsecase) GetAllTests(req request.QuizParamRequest) (*[]dto.QuizResponse, *dto.Meta, error) {
	quizzes, meta, err := tc.testRepo.GetAllTests(req)
	if err != nil {
		return quizzes, meta, err
	}

	return quizzes, meta, err
}
