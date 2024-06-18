package materialAdmin

import (
	"fmt"
	"math"

	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/response"
)

func (u *usecase) MaterialAdminAll(req request.ParamMaterialAdminAll) ([]*response.MaterialAdminAllResponse, int, int, error) {
	limit := req.Limit
	offset := 0

	if req.Page > 0 {
		offset = (req.Page - 1) * limit
	}

	var subModuleID, subjectID int

	if req.SubModuleName != "" {
		subModule, err := u.subModulRepo.FindOneBy(map[string]interface{}{
			"name": req.SubModuleName,
		})
		if err != nil {
			return nil, 0, 0, fmt.Errorf("failed to find subModule: %v", err)
		}
		subModuleID = subModule.ID
	}

	if req.SubjectName != "" {
		subject, err := u.subjectRepo.FindOneBy(map[string]interface{}{
			"name": req.SubjectName,
		})
		if err != nil {
			return nil, 0, 0, fmt.Errorf("failed to find subject: %v", err)
		}
		subjectID = subject.ID
	}

	countMaterialAll, err := u.materialAdminLocalRepo.CountMaterialAdminALL(req.SearchKeywords, subjectID, subModuleID)
	totalRows := countMaterialAll
	totalPages := int(math.Ceil(float64(countMaterialAll) / float64(req.Limit)))
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to countMaterialAll: %v", err)
	}

	materialAll, err := u.materialAdminLocalRepo.GetMaterialAdminAll(req.SearchKeywords, subjectID, subModuleID, limit, offset)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to find materialAll: %v", err)
	}

	result := materialAll

	return result, totalRows, totalPages, nil

}
