package module

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (mod *moduleUsecase) GetListModules() ([]dto.ModuleResponse, error) {
	fmt.Println("GetListModules usecase")
	moduleList, err := mod.moduleRepo.GetListModules()
	if err != nil {
		return nil, err
	}
	moduleTestList, err := mod.moduleRepo.GetListSubjectQuiz()
	if err != nil {
		return nil, err
	}

	var moduleResponses []dto.ModuleResponse
	for _, module := range moduleList {
		moduleResponse := dto.ModuleResponse{
			ID:          module.ID,
			Name:        module.Name,
			Description: module.Description,
		}
		for _, subject := range module.Subjects {
			subjectResponse := struct {
				ID   string `json:"ID" form:"ID"`
				Name string `json:"Name" form:"Name"`
			}{
				ID:   subject.ID,
				Name: subject.Name,
			}
			for _, moduleTest := range moduleTestList {
				if moduleTest.SubjectID == subject.ID {
					moduleTestResponse := struct {
						ID   string `json:"ID" form:"ID"`
						Name string `json:"Name" form:"Name"`
					}{
						ID:   moduleTest.ID,
						Name: subject.Name,
					}

					moduleResponse.Exercise = append(moduleResponse.Exercise, moduleTestResponse)
				}
			}
			moduleResponse.Subject = append(moduleResponse.Subject, subjectResponse)
		}
		moduleResponses = append(moduleResponses, moduleResponse)
	}
	return moduleResponses, nil
}

func (mod *moduleUsecase) GetSubjectsBySubmoduleID(submoduleID string) ([]dto.SubjectResponse, error) {
	fmt.Println("GetSubjectsBySubmoduleID usecase")

	subjects, err := mod.moduleRepo.GetSubjectsBySubmoduleID(submoduleID)
	if err != nil {
		return nil, err
	}

	var subjectResponses []dto.SubjectResponse

	for _, subject := range subjects {
		subjectResponse := dto.SubjectResponse{
			ID:   subject.ID,
			Name: subject.Name,
		}
		for _, subSubject := range subject.SubSubject {
			subSubjectResponse := dto.SubSubject{
				ID:      subSubject.ID,
				Name:    subSubject.Name,
				Content: subSubject.Content,
			}
			subjectResponse.SubSubject = append(subjectResponse.SubSubject, subSubjectResponse)
		}

		subjectResponses = append(subjectResponses, subjectResponse)

	}

	return subjectResponses, nil
}

func (mod *moduleUsecase) GetQuestionsByTestTypeID(testTypeID string) (dto.ExerciseResponse, error) {
	fmt.Println("GetQuestionsByTestTypeID usecase")

	var moduleTest datastruct.SubjectTestTypeQuizzes

	moduleTest, err := mod.moduleRepo.GetTestByTestTypeID(testTypeID)
	if err != nil {
		return dto.ExerciseResponse{}, err
	}

	var moduleTestResponses dto.ExerciseResponse

	moduleTestResponses.ID = moduleTest.ID
	moduleTestResponses.Name = moduleTest.Title
	moduleTestResponses.Description = moduleTest.Description
	moduleTestResponses.Explanation = moduleTest.Explanation
	moduleTestResponses.Status = moduleTest.Status
	moduleTestResponses.Time = moduleTest.Time
	moduleTestResponses.QuestionTotal = len(moduleTest.QuestionQuizzes)

	for i, question := range moduleTest.QuestionQuizzes {
		questionResponse := dto.Question{
			ID:       question.ID,
			Number:   i + 1,
			Status:   question.Status,
			Mark:     1,
			Flag:     false,
			Question: question.Content,
			Image:    question.Image,
		}
		for _, answer := range question.ChoiceQuizzes {
			answerResponse := dto.Answer{
				ID:        answer.ID,
				Content:   answer.Content,
				IsCorrect: answer.IsCorrect,
				Weight:    answer.Weight,
				IsChosen:  false,
			}
			questionResponse.Answers = append(questionResponse.Answers, answerResponse)
		}
		moduleTestResponses.Questions = append(moduleTestResponses.Questions, questionResponse)
	}

	return moduleTestResponses, nil
}

func (mod *moduleUsecase) GetTestReview(moduleTestID string) (dto.ReviewResultResponse, error) {
	fmt.Println("GetQuestionsByTestTypeID usecase")

	var moduleTest datastruct.SubjectTestTypeQuizzes

	answeredChoices, err := mod.moduleRepo.GetSubmittedReviewByID(moduleTestID)
	if err != nil {
		return dto.ReviewResultResponse{}, err
	}

	moduleTest, err = mod.moduleRepo.GetTestReview(answeredChoices.TestTypeQuizID)
	if err != nil {
		return dto.ReviewResultResponse{}, err
	}

	var moduleTestResponses dto.ReviewResultResponse

	moduleTestResponses.ID = moduleTest.ID
	moduleTestResponses.Name = moduleTest.Title
	moduleTestResponses.Description = moduleTest.Description
	moduleTestResponses.Explanation = moduleTest.Explanation
	moduleTestResponses.Status = moduleTest.Status
	moduleTestResponses.Time = string(moduleTest.Time)
	moduleTestResponses.QuestionTotal = len(moduleTest.QuestionQuizzes)
	moduleTestResponses.Grade = answeredChoices.GradeQuiz.Score

	for i, question := range moduleTest.QuestionQuizzes {
		questionResponse := dto.Question{
			ID:       question.ID,
			Number:   i + 1,
			Status:   question.Status,
			Mark:     1,
			Flag:     false,
			Question: question.Content,
			Image:    question.Image,
		}

		// Create a map to store submitted choices for the current question.
		submittedChoices := make(map[string]bool)
		for _, userAnswer := range answeredChoices.UserSubmittedAnswerQuizzes {
			if userAnswer.QuestionQuizID == question.ID {
				submittedChoices[userAnswer.ChoiceQuizID] = true
			}
		}

		// Iterate over choices and build answer responses.
		for _, answer := range question.ChoiceQuizzes {
			answerResponse := dto.Answer{
				ID:        answer.ID,
				Content:   answer.Content,
				IsCorrect: answer.IsCorrect,
				Weight:    answer.Weight,
				IsChosen:  submittedChoices[answer.ID], // Check if the answer is chosen by the user.
			}
			if answerResponse.IsChosen && !answerResponse.IsCorrect {
				questionResponse.Status = "Wrong"
			}
			if answerResponse.IsChosen && answerResponse.IsCorrect {
				questionResponse.Status = "Correct"
			}
			questionResponse.Answers = append(questionResponse.Answers, answerResponse)
		}

		moduleTestResponses.Questions = append(moduleTestResponses.Questions, questionResponse)
	}

	return moduleTestResponses, nil
}

func (mod *moduleUsecase) PostSubmittedTest(testTypeID string, submittedQuiz dto.UserSubmittedQuizRequest) error {
	// fmt.Println("PostSubmittedTest usecase")

	submitedTest := datastruct.UserTestSubmissionQuizzes{
		ID:             uuid.New().String(),
		CreatedAt:      time.Now(),
		UserID:         submittedQuiz.UserID,
		TestTypeQuizID: testTypeID,
		SubmissionTIme: time.Now(),
	}
	submissionID, err := mod.moduleRepo.AddQuizSubmission(submitedTest)
	if err != nil {
		return err
	}
	if submissionID == "" {
		return fmt.Errorf("Failed to submit test")
	}

	submittedAnswers := make([]datastruct.UserSubmittedAnswerQuizzes, 0)
	for _, pair := range submittedQuiz.PairOfUserAnswer {
		submittedAnswers = append(submittedAnswers, datastruct.UserSubmittedAnswerQuizzes{
			ID:                       uuid.New().String(),
			CreatedAt:                time.Now(),
			UserTestSubmissionQuizID: submissionID,
			QuestionQuizID:           pair.QuestionQuizID,
			ChoiceQuizID:             pair.ChoiceQuizID,
		})

	}

	quizGradeResult := datastruct.GradeQuizzes{
		ID:                       uuid.New().String(),
		UserTestSubmissionQuizID: submissionID,
		UserID:                   submittedQuiz.UserID,
		TestTypeQuizID:           testTypeID,
		GradingTime:              time.Now(),
	}
	err = mod.moduleRepo.AddGrade(quizGradeResult)
	if err != nil {
		return fmt.Errorf("Failed grading test:")
	}

	err = mod.moduleRepo.PostSubmittedQuizAnswer(submissionID, submittedAnswers)
	if err != nil {
		return err
	}

	return nil
}
