package class

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
)

func (u *usecase) FindAll(req request.Class) ([]*response.Class, error) {
	user, err := u.userRepo.FindOneBy(map[string]interface{}{
		"id":      req.AuthenticatedUser,
		"role_id": helper.PembeliPaketBimbel,
	})
	if user == nil {
		return nil, helper.ErrAccessDenied
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}
	return nil, nil
}
