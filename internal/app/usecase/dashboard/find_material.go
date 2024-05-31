package dashboard

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) FindMaterial(req request.ParamDashboard) ([]*response.MaterialDashboard, error) {
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

	subModules, err := u.subModuleRepo.FindBy(map[string]interface{}{
		"id": subModuleIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find sub modules: %v", err)
	}

	subModuleMap := make(map[int]*model.SubModule)
	for _, subModule := range subModules {
		subModuleMap[subModule.ID] = subModule
	}

	moduleMap := make(map[int]*model.Module)
	for _, module := range modules {
		moduleMap[module.ID] = module
	}

	res := make([]*response.MaterialDashboard, 0)
	for _, subject := range subjects {
		subModule := subModuleMap[subject.SubModuleID]
		if subModule == nil {
			return nil, fmt.Errorf("module not found for subModuleID: %d", subject.SubModuleID)
		}

		module := moduleMap[subModule.ModuleID]
		if module == nil {
			return nil, fmt.Errorf("module not found for moduleID: %d", subModule.ModuleID)
		}

		res = append(res, &response.MaterialDashboard{
			UUID: subject.UUID,
			Module: response.Module{
				UUID: module.UUID,
				Name: module.Name,
			},
			SubModule: response.SubModule{
				UUID: subModule.UUID,
				Name: subModule.Name,
			},
			Subject: subject.Name,
		})
	}

	return res, nil
}
