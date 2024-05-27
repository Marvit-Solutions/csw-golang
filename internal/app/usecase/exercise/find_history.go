package exercise

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
)

func (u *usecase) FindHistory(req request.ExerciseHistory) ([]*response.ExerciseHistory, error) {
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

	exerciseSubmissions, err := u.exerciseSubmissionRepo.FindBy(map[string]interface{}{
		"exercise_id": exercise.ID,
		"user_id":     req.AuthenticatedUser,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find exercise submissions: %v", err)
	}

	res := make([]*response.ExerciseHistory, 0)
	for _, exerciseSubmission := range exerciseSubmissions {
		res = append(res, &response.ExerciseHistory{
			SubmissionUUID: exerciseSubmission.UUID,
			Score:          exerciseSubmission.Score,
		})
	}

	return res, nil
}
