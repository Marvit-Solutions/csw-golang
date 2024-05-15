package exercise

import (
	"fmt"
	"strings"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
)

func (u *usecase) FindAll(req request.ParamExercise) ([]*response.ExerciseResponse, error) {

	moduleMap := map[string]int{
		"skd":        helper.SKDModuleID,
		"matematika": helper.MatematikaModuleID,
	}

	moduleID := []int{}
	moduleReq := strings.ToLower(req.Module)

	if id, exists := moduleMap[moduleReq]; exists {
		moduleID = append(moduleID, id)
	} else {
		moduleID = append(moduleID,
			helper.SKDModuleID,
			helper.MatematikaModuleID)
	}

	testTypeMap := map[string]int{
		"latihan-soal": helper.LatihanSoalTestTypeID,
		"pretest":      helper.PretestTestTypeID,
		"posttest":     helper.PosttestTestTypeID,
		"paket-soal":   helper.PaketSoalTestTypeID,
		"try-out":      helper.TryOutTestTypeID,
	}

	testTypeID := []int{}
	testTypeReq := strings.ToLower(req.TestType)

	if id, exists := testTypeMap[testTypeReq]; exists {
		testTypeID = append(testTypeID, id)
	} else {
		testTypeID = append(testTypeID,
			helper.LatihanSoalTestTypeID,
			helper.PretestTestTypeID,
			helper.PosttestTestTypeID,
			helper.PaketSoalTestTypeID,
			helper.TryOutTestTypeID)
	}

	exercises, err := u.exerciseRepo.FindBy(map[string]interface{}{
		"module_id":  moduleID,
		"deleted_at": nil,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find exercises: %v", err)
	}

	module, err := u.moduleRepo.FindOneBy(map[string]interface{}{
		"id": moduleID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find module: %v", err)
	}

	testType, err := u.testTypeRepo.FindOneBy(map[string]interface{}{
		"id": testTypeID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find test type: %v", err)
	}

	results := make([]*response.ExerciseResponse, 0)
	for _, exercise := range exercises {
		results = append(results, &response.ExerciseResponse{
			UUID:          exercise.UUID,
			TestType:      testType.Name,
			ModuleName:    module.Name,
			Title:         exercise.Title,
			TotalQuestion: exercise.TotalQuestion,
			Time:          exercise.Time,
			Description:   exercise.Description,
		})
	}

	return results, nil
}
