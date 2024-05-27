package dashboard

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) FindMaterial(req request.Dashboard) ([]*response.MaterialDashboard, error) {
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

	subjects, err := u.subjectRepo.FindBy(map[string]interface{}{}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find subjects: %v", err)
	}

	subModuleIDs := make([]int, 0, len(subjects))
	for _, subject := range subjects {
		subModuleIDs = append(subModuleIDs, subject.SubModuleID)
	}

	modules, err := u.moduleRepo.FindBy(map[string]interface{}{
		"id": subModuleIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find modules: %v", err)
	}

	moduleMap := make(map[int]*model.Module)
	for _, module := range modules {
		moduleMap[module.ID] = module
	}

	skdIDs := []int{helper.TWK, helper.TIU, helper.TKP}
	isSKD := func(id int) bool {
		for _, SKD := range skdIDs {
			if id == SKD {
				return true
			}
		}
		return false
	}

	res := make([]*response.MaterialDashboard, 0)
	for _, subject := range subjects {
		var moduleName string
		if isSKD(subject.SubModuleID) {
			moduleName = moduleMap[helper.SKDModuleID].Name
		} else {
			moduleName = moduleMap[helper.MatematikaModuleID].Name
		}
		res = append(res, &response.MaterialDashboard{
			UUID:    subject.UUID,
			Module:  moduleName,
			Subject: subject.Name,
		})
	}

	return res, nil
}
