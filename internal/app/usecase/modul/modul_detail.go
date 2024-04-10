package modul

import (
	"fmt"
	"sort"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
)

func (u *usecase) ModulDetail(req request.ParamModulDetail) (*response.ModulDetailResponse, error) {
	subModule, err := u.subModuleRepo.FindOneBy(map[string]interface{}{
		"uuid": req.UUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find sub module: %v", err)
	}

	module, err := u.moduleRepo.FindOneBy(map[string]interface{}{
		"id": subModule.ModuleID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find module: %v", err)
	}

	subjects, err := u.subjectRepo.FindBy(map[string]interface{}{
		"sub_module_id": subModule.ID,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find subjects: %v", err)
	}

	result := &response.ModulDetailResponse{
		UUID:          subModule.UUID,
		SubModuleName: subModule.Name,
		Description:   subModule.Description,
		ModuleName:    module.Name,
	}

	resp := make([]*response.SubjectResponse, 0)
	for _, subject := range subjects {
		resp = append(resp, &response.SubjectResponse{
			ID:   subject.ID,
			UUID: subject.UUID,
			Name: subject.Name,
		})
	}

	result.Subjects = resp

	sort.Slice(result.Subjects, func(i, j int) bool {
		return result.Subjects[i].ID < result.Subjects[j].ID
	})

	return result, nil
}
