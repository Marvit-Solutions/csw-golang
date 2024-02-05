package test

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
	"github.com/google/uuid"
)

func (e exerciseTestUsecase) AddExerciseTestUsecase(request dto.SubjectTestTypeExercisesRequest) (dto.SubjectTestTypeExercisesResponse, error) {

	var questionsArr []datastruct.QuestionExercises
	idTestType := uuid.NewString()
	for _, request := range request.QuestionExercises {
		var choiceExercisesArray []datastruct.ChoiceExercises
		for _, request1 := range request.ChoiceExercises {
			choiceExercises := datastruct.ChoiceExercises{
				ID:        uuid.NewString(),
				Content:   request1.Content,
				IsCorrect: request1.IsCorrect,
				Weight:    request1.Weight,
			}

			choiceExercisesArray = append(choiceExercisesArray, choiceExercises)
		}

		exercises := datastruct.QuestionExercises{
			ID:                 uuid.New().String(),
			Content:            request.Content,
			Weight:             request.Weight,
			TestTypeExerciseID: idTestType,
			ChoiceExercises:    choiceExercisesArray,
		}
		questionsArr = append(questionsArr, exercises)
	}

	exercises := datastruct.SubjectTestTypeExercises{
		ID:                idTestType,
		TestType:          request.TestType,
		QuestionExercises: questionsArr,
	}

	test, err := e.exerciseTestRepo.AddExerciseTest(exercises)
	if err != nil {
		return dto.SubjectTestTypeExercisesResponse{}, err
	}

	response := dto.SubjectTestTypeExercisesResponse{
		Id:       test.ID,
		TestType: test.TestType,
	}
	return response, nil
}

func (e exerciseTestUsecase) GetAllExerciseTestUsecase() ([]dto.SubjectTestTypeExercisesResponse, error) {

	usecase, err := e.exerciseTestRepo.GetAllExerciseTest()
	if err != nil {
		return []dto.SubjectTestTypeExercisesResponse{}, err
	}

	var testType []dto.SubjectTestTypeExercisesResponse
	for _, use := range usecase {
		response := dto.SubjectTestTypeExercisesResponse{
			Id:       use.ID,
			TestType: use.TestType,
		}
		testType = append(testType, response)
	}

	return testType, nil

}

func (e exerciseTestUsecase) GetExerciseTestUsecase(id string) (dto.SubjectTestTypeExercisesResponse, error) {
	test, err := e.exerciseTestRepo.GetExerciseTest(id)
	if err != nil {
		return dto.SubjectTestTypeExercisesResponse{}, err
	}

	var questionsArr []dto.QuestionExercisesRequest
	for _, exe := range test.QuestionExercises {
		var choice []dto.ChoiceExercisesRequest
		for _, ch := range exe.ChoiceExercises {
			request := dto.ChoiceExercisesRequest{
				Content:   ch.Content,
				IsCorrect: ch.IsCorrect,
				Weight:    ch.Weight,
			}
			choice = append(choice, request)
		}

		request := dto.QuestionExercisesRequest{
			Content:         exe.Content,
			Weight:          exe.Weight,
			ChoiceExercises: choice,
		}
		questionsArr = append(questionsArr, request)
	}

	response := dto.SubjectTestTypeExercisesResponse{
		Id:                test.ID,
		TestType:          test.TestType,
		QuestionExercises: questionsArr,
	}
	return response, nil
}
