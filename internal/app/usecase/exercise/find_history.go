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

	return nil, nil
}
