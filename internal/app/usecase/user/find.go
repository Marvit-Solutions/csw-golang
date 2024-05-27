package user

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) Find(req request.User) (*response.User, error) {
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

	userDetail, err := u.userDetailRepo.FindOneBy(map[string]interface{}{
		"user_id": user.ID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find user detail: %v", err)
	}

	var media *model.Media
	if userDetail.MediaID != nil {
		media, err = u.mediaRepo.FindOneBy(map[string]interface{}{
			"id": userDetail.MediaID,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to find media: %v", err)
		}
	}

	classUser, err := u.classUserRepo.FindOneBy(map[string]interface{}{
		"id": userDetail.ClassUserID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find class user: %v", err)
	}

	res := &response.User{
		UUID:        user.UUID,
		Email:       user.Email,
		Name:        userDetail.Name,
		PhoneNumber: userDetail.PhoneNumber,
		Province:    userDetail.Province,
		Regency:     userDetail.Regency,
		District:    userDetail.District,
		Media:       helper.MultiResImages(media),
		Class:       classUser.Name,
	}

	return res, nil
}
