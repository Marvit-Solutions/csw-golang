package tests

import "csw-golang/internal/domain/entity/datastruct"

func (tc *testUsecase) GetAllTests() ([]datastruct.Modules, error) {
	ListPlan, err := tc.testRepo.GetAllTests()
	if err != nil {
		return ListPlan, err
	}

	return ListPlan, err
}
