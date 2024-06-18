package subjectAdmin

import (
	"fmt"
	"math"

	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/response"
)

func (u *usecase) SubjectAdminAll(req request.ParamSubjectAdminAll) ([]*response.SubjectAdminAllResponse, int, int, error) {
	limit := req.Limit
	offset := 0

	if req.Page > 0 {
		offset = (req.Page - 1) * limit
	}

	var subModuleID int

	if req.SubModuleName != "" {
		subModule, err := u.subModulRepo.FindOneBy(map[string]interface{}{
			"name": req.SubModuleName,
		})
		if err != nil {
			return nil, 0, 0, fmt.Errorf("failed to find subModule: %v", err)
		}
		subModuleID = subModule.ID
	}

	countSubjectAll, err := u.subjectAdminLocalRepo.CountSubjectAdminALL(req.SearchKeywords, subModuleID)
	totalRows := countSubjectAll
	totalPages := int(math.Ceil(float64(countSubjectAll) / float64(req.Limit)))
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to countSubjectAll: %v", err)
	}

	subjectAll, err := u.subjectAdminLocalRepo.GetSubjectAdminAll(req.SearchKeywords, subModuleID, limit, offset)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to find subjectAll: %v", err)
	}

	result := subjectAll

	return result, totalRows, totalPages, nil

}
