package tests

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"
	"csw-golang/internal/domain/helper/response"
)

func (tc *testUsecase) GetAllTests(req request.QuizRequest) (*[]dto.QuizResponse, *response.Meta, error) {
	quizzes, meta, err := tc.testRepo.GetAllTests(req)
	if err != nil {
		return quizzes, meta, err
	}

	return quizzes, meta, err
}
