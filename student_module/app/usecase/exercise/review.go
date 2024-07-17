package exercise

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/response"
)

func (u *usecase) Review(req request.ExerciseReview) (*response.ExerciseReview, error) {
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

	exerciseSubmissionModule, err := u.exerciseSubmissionsModuleRepo.FindOneBy(map[string]interface{}{
		"uuid": req.SubmissionUUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find exercise submission: %v", err)
	}

	exercise, err := u.exerciseRepo.FindOneBy(map[string]interface{}{
		"id": exerciseSubmissionModule.ExerciseID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find exercise : %v", err)
	}

	module, err := u.moduleRepo.FindOneBy(map[string]interface{}{
		"id": exercise.ModuleID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find module: %v", err)
	}

	totalQuestion := u.exerciseQuestionRepo.Count(map[string]interface{}{
		"exercise_id": exerciseSubmissionModule.ExerciseID,
	})

	exerciseQuestions, err := u.exerciseQuestionRepo.FindBy(map[string]interface{}{
		"exercise_id": exerciseSubmissionModule.ExerciseID,
	}, 0, totalQuestion)
	if err != nil {
		return nil, fmt.Errorf("failed to find exercise questions: %v", err)
	}

	questionIDs := make([]int, 0, len(exerciseQuestions))
	for _, question := range exerciseQuestions {
		questionIDs = append(questionIDs, question.ID)
	}

	var perfectScore int
	for _, question := range exerciseQuestions {
		perfectScore += question.Score
	}

	exerciseSubmissionsSubModule, err := u.exerciseSubmissionsSubModuleRepo.FindBy(map[string]interface{}{
		"submissions_module_id": exerciseSubmissionModule.ID,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find exercise submission: %v", err)
	}

	exerciseSubmissionSubModuleIDs := make([]int, 0, len(exerciseSubmissionsSubModule))
	for _, exerciseSubmissionSubModule := range exerciseSubmissionsSubModule {
		exerciseSubmissionSubModuleIDs = append(exerciseSubmissionSubModuleIDs, exerciseSubmissionSubModule.ID)
	}

	exerciseAnswers, err := u.exerciseAnswerRepo.FindBy(map[string]interface{}{
		"submission_sub_module_id": exerciseSubmissionSubModuleIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find exercise answers: %v", err)
	}

	userAnswerIDs := make(map[int]struct{})
	for _, answer := range exerciseAnswers {
		if answer.ChoiceID != nil {
			userAnswerIDs[*answer.ChoiceID] = struct{}{}
		}
	}

	choices, err := u.exerciseChoiceRepo.FindBy(map[string]interface{}{
		"question_id": questionIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find choices: %v", err)
	}

	filteredChoices := make(map[int]*model.ExerciseChoice)
	for _, choice := range choices {
		if _, ok := userAnswerIDs[choice.ID]; ok {
			filteredChoices[choice.ID] = choice
		}
	}

	choiceResMap := make(map[int][]*response.ChoiceReview)
	for _, choice := range choices {
		isChosen := false
		if _, ok := filteredChoices[choice.ID]; ok {
			isChosen = true
		}
		choiceResMap[choice.QuestionID] = append(choiceResMap[choice.QuestionID], &response.ChoiceReview{
			ID:         choice.ID,
			UUID:       choice.UUID,
			Content:    choice.Content,
			QuestionID: choice.QuestionID,
			IsChoose:   isChosen,
			IsCorrect:  choice.IsCorrect,
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

	questions, err := u.exerciseQuestionRepo.FindBy(map[string]interface{}{
		"exercise_id": exerciseSubmissionModule.ExerciseID,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find questions: %v", err)
	}

	questionsRes := make([]*response.QuestionReview, len(questions))
	for i, question := range questions {
		var userAnswer int
		var rigthAnswer int
		var rigthAnswerText string
		var status response.QuestionReviewItemStatus

		for _, choice := range choiceResMap[question.ID] {
			if choice.IsChoose {
				userAnswer = choice.ID
			}
			if choice.IsCorrect {
				rigthAnswer = choice.ID
				rigthAnswerText = choice.Content
			}

		}

		if userAnswer == 0 {
			status = response.BelumDiJawab
		} else {
			status = response.SudahDiJawab
		}
		questionsRes[i] = &response.QuestionReview{
			ID:              question.ID,
			UUID:            question.UUID,
			Content:         question.Content,
			Status:          string(status),
			Explanation:     question.Explanation,
			Score:           question.Score,
			UserAnswer:      userAnswer,
			RightAnswer:     rigthAnswer,
			RightAnswerText: rigthAnswerText,
			QuestionMedia:   questionMediaMap[question.ID],
			Choices:         choiceResMap[question.ID],
		}
	}

	attempt := u.exerciseSubmissionsModuleRepo.Count(map[string]interface{}{
		"user_id":     user.ID,
		"exercise_id": exerciseSubmissionModule.ExerciseID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find exercise submission: %v", err)
	}

	res := &response.ExerciseReview{
		ID:            exerciseSubmissionModule.ID,
		UUID:          exerciseSubmissionModule.UUID,
		Topic:         exercise.Title,
		Modul:         module.Name,
		StartedAt:     helper.ConvertToIndonesianFormat(exerciseSubmissionModule.StartedAt),
		FinishedAt:    helper.ConvertToIndonesianFormat(exerciseSubmissionModule.FinishedAt),
		TimeRequired:  helper.ConvertDurationToIndonesian(exerciseSubmissionModule.TimeRequired),
		Attempt:       attempt,
		RightAnswer:   exerciseSubmissionModule.RightAnswer,
		TotalQuestion: totalQuestion,
		Score:         exerciseSubmissionModule.Score,
		PerfectScore:  perfectScore,
		Questions:     questionsRes,
	}

	return res, nil
}
