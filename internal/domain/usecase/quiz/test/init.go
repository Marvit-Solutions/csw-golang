package tests

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"
	"csw-golang/internal/domain/helper/response"
	tr "csw-golang/internal/domain/repository/quiz/test"
)

type TestUsecase interface {
	GetAllTests(req request.QuizParamRequest) (*[]dto.QuizResponse, *response.Meta, error)
}

type testUsecase struct {
	testRepo tr.TestRepo
}

func New(testRepo tr.TestRepo) *testUsecase {
	return &testUsecase{
		testRepo,
	}
}
