package exercise

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
)

func (u *usecase) Review(req request.ExerciseReview) ([]*response.ExerciseReview, error) {

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

	var perfectScore int
	for _, question := range exerciseQuestions {
		perfectScore += question.Score
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
	}

	fmt.Println(helper.Debug(res))

	return nil, nil
}
