package subjectAdmin

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/response"
)

func (u *usecase) SubjectAdminUpdateDetail(req request.ParamSubjectAdminUpdateDetail) (*response.SubjectAdminUpdateDetailResponse, error) {
	subject, err := u.subjectRepo.FindOneBy(map[string]interface{}{
		"uuid": req.UUID,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to find subject: %v", err)
	}

	subModule, err := u.subModulRepo.FindOneBy(map[string]interface{}{
		"id": subject.SubModuleID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find subModule: %v", err)
	}

	result := &response.SubjectAdminUpdateDetailResponse{
		UUID:          subject.UUID,
		SubModuleName: subModule.Name,
		Name:          subject.Name,
	}

	return result, nil
}
