package tests

import (
	tc "csw-golang/internal/domain/usecase/quiz/test"
)

type TestHandler struct {
	testUsecase tc.TestUsecase
}

func New(testUsecase tc.TestUsecase) *TestHandler {
	return &TestHandler{
		testUsecase,
	}
}
