package tests

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"
)

func (tc *testUsecase) GetAllTests(req request.QuizParamRequest) (*[]dto.QuizResponse, error) {
	quizzes, err := tc.testRepo.GetAllTests(req)
	if err != nil {
		return quizzes, err
	}

	return quizzes, err
}
