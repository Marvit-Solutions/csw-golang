package modul

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) ModulAll() ([]*response.ModulResponse, error) {
	subModules, err := u.subModuleRepo.FindBy(map[string]interface{}{}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find sub modules: %v", err)
	}

	moduleIDs := make([]int, len(subModules))
	for i, subModule := range subModules {
		moduleIDs[i] = subModule.ModuleID
	}

	modules, err := u.moduleRepo.FindBy(map[string]interface{}{
		"id": moduleIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find modules: %v", err)
	}

	ModuleMap := make(map[int]*model.Module)
	for _, module := range modules {
		ModuleMap[module.ID] = module
	}

	results := make([]*response.ModulResponse, 0)
	for _, subModule := range subModules {
		results = append(results, &response.ModulResponse{
			UUID:          subModule.UUID,
			SubModuleName: subModule.Name,
			ModuleName:    ModuleMap[subModule.ModuleID].Name,
			Description:   subModule.Description,
		})
	}

	return results, nil
}
