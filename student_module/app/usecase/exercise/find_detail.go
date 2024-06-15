package exercise

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/response"
)

func (u *usecase) FindDetail(req request.ExerciseDetailRequest) (*response.ExerciseDetail, error) {
	user, err := u.userRepo.FindOneBy(map[string]interface{}{
		"id":      req.AuthenticatedUser,
		"role_id": helper.PembeliPaketBimbel,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}
	if user == nil {
		return nil, helper.ErrAccessDenied
	}

	exercise, err := u.exerciseRepo.FindOneBy(map[string]interface{}{
		"uuid": req.ExerciseUUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find exercise: %v", err)
	}

	userAttempt := u.exerciseSubmissionRepo.Count(map[string]interface{}{
		"exercise_id": exercise.ID,
		"user_id":     req.AuthenticatedUser,
	})

	testType, err := u.testTypeRepo.FindOneBy(map[string]interface{}{
		"id": exercise.TestTypeID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find test type: %v", err)
	}

	module, err := u.moduleRepo.FindOneBy(map[string]interface{}{
		"id": exercise.ModuleID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find module: %v", err)
	}

	questions, err := u.exerciseQuestionRepo.FindBy(map[string]interface{}{
		"exercise_id": exercise.ID,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find questions: %v", err)
	}

	questionIDs := make([]int, len(questions))
	for i, question := range questions {
		questionIDs[i] = question.ID
	}

	choices, err := u.exerciseChoiceRepo.FindBy(map[string]interface{}{
		"question_id": questionIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find choices: %v", err)
	}

	choiceResMap := make(map[int][]*response.Choice)
	for _, choice := range choices {
		choiceResMap[choice.QuestionID] = append(choiceResMap[choice.QuestionID], &response.Choice{
			UUID:       choice.UUID,
			Content:    choice.Content,
			QuestionID: choice.QuestionID,
		})
	}

	questionMedias, err := u.exerciseQuestionMediaRepo.FindBy(map[string]interface{}{
		"exercise_question_id": questionIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find question medias: %v", err)
	}

	mediaIDs := make([]int, len(questionMedias))
	for i, questionMedia := range questionMedias {
		mediaIDs[i] = questionMedia.MediaID
	}

	medias, err := u.mediaRepo.FindBy(map[string]interface{}{
		"id": mediaIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find medias: %v", err)
	}

	mediaMaps := make(map[int]*model.Media)
	for _, media := range medias {
		mediaMaps[media.ID] = media
	}

	questionMediaMap := make(map[int][]*response.QuestionMedia)
	for _, questionMedia := range questionMedias {
		questionMediaMap[questionMedia.ExerciseQuestionID] = append(questionMediaMap[questionMedia.ExerciseQuestionID], &response.QuestionMedia{
			Index: questionMedia.Index,
			Media: helper.MultiResImages(mediaMaps[questionMedia.MediaID]),
		})
	}

	questionsRes := make([]*response.Question, len(questions))
	for i, question := range questions {
		questionsRes[i] = &response.Question{
			UUID:          question.UUID,
			Content:       question.Content,
			Score:         question.Score,
			QuestionMedia: questionMediaMap[question.ID],
			Choices:       choiceResMap[question.ID],
		}
	}

	result := &response.ExerciseDetail{
		UUID:        exercise.UUID,
		TestType:    testType.Name,
		ModuleName:  module.Name,
		Title:       exercise.Title,
		Attempt:     exercise.Attempt,
		UserAttempt: userAttempt,
		Time:        exercise.Time,
		Description: exercise.Description,
		Questions:   questionsRes,
	}

	return result, nil
}
