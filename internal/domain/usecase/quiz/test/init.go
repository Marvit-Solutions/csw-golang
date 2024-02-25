package tests

import (
	"csw-golang/internal/domain/entity/dto"
	tr "csw-golang/internal/domain/repository/quiz/test"
)

type TestUsecase interface {
	GetAllTests() ([]dto.QuizResponse, error)
}

type testUsecase struct {
	testRepo tr.TestRepo
}

func New(testRepo tr.TestRepo) *testUsecase {
	return &testUsecase{
		testRepo,
	}
}
