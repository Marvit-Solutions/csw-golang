package tests

import (
	"csw-golang/internal/domain/entity/datastruct"
	tr "csw-golang/internal/domain/repository/quiz/test"
)

type TestUsecase interface {
	GetAllTests() ([]datastruct.Modules, error)
}

type testUsecase struct {
	testRepo tr.TestRepo
}

func New(testRepo tr.TestRepo) *testUsecase {
	return &testUsecase{
		testRepo,
	}
}
