package exercise

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
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

	exerciseSubmission, err := u.exerciseSubmissionRepo.FindOneBy(map[string]interface{}{
		"uuid": req.SubmissionUUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find exercise submission: %v", err)
	}

	totalQuestion := u.exerciseQuestionRepo.Count(map[string]interface{}{
		"exercise_id": exerciseSubmission.ExerciseID,
	})

	exerciseQuestions, err := u.exerciseQuestionRepo.FindBy(map[string]interface{}{
		"exercise_id": exerciseSubmission.ExerciseID,
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

	exerciseAnswers, err := u.exerciseAnswerRepo.FindBy(map[string]interface{}{
		"submission_id": exerciseSubmission.ID,
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
		"exercise_id": exerciseSubmission.ExerciseID,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find questions: %v", err)
	}

	questionsRes := make([]*response.QuestionReview, len(questions))
	for i, question := range questions {
		questionsRes[i] = &response.QuestionReview{
			UUID:          question.UUID,
			Content:       question.Content,
			Score:         question.Score,
			QuestionMedia: questionMediaMap[question.ID],
			Choices:       choiceResMap[question.ID],
		}
	}

	res := &response.ExerciseReview{
		UUID:          exerciseSubmission.UUID,
		StartedAt:     helper.ConvertToIndonesianFormat(exerciseSubmission.StartedAt),
		FinishedAt:    helper.ConvertToIndonesianFormat(exerciseSubmission.FinishedAt),
		TimeRequired:  helper.ConvertDurationToIndonesian(exerciseSubmission.TimeRequired),
		RightAnswer:   exerciseSubmission.RightAnswer,
		TotalQuestion: totalQuestion,
		Score:         exerciseSubmission.Score,
		PerfectScore:  perfectScore,
		Questions:     questionsRes,
	}

	return res, nil
}
