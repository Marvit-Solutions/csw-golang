package materialAdmin

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/response"
)

func (u *usecase) MaterialAdminUpdateDetail(req request.ParamMaterialAdminUpdateDetail) (*response.MaterialAdminUpdateDetailResponse, error) {
	subSubject, err := u.subSubjectRepo.FindOneBy(map[string]interface{}{
		"uuid": req.UUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find subSubject: %v", err)
	}

	subject, err := u.subjectRepo.FindOneBy(map[string]interface{}{
		"id": subSubject.SubjectID,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to find subject: %v", err)
	}

	result := &response.MaterialAdminUpdateDetailResponse{
		UUID:        subSubject.UUID,
		SubjectName: subject.Name,
		Name:        subSubject.Name,
		Content:     subSubject.Content,
	}

	return result, nil
}
