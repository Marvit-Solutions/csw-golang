package tests

import "csw-golang/internal/domain/entity/dto"

func (tc *testUsecase) GetAllTests() ([]dto.QuizResponse, error) {
	quizzes, err := tc.testRepo.GetAllTests()
	if err != nil {
		return quizzes, err
	}

	return quizzes, err
}
