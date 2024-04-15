package module

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
)

func (u *usecase) MaterialAll(req request.ParamModule) (*response.MaterialResponse, error) {
	subject, err := u.subjectRepo.FindOneBy(map[string]interface{}{
		"uuid": req.SubjectUUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find subject: %v", err)
	}

	subSubjects, err := u.subSubjectRepo.FindBy(map[string]interface{}{
		"subject_id": subject.ID,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find sub subjects: %v", err)
	}

	subModule, err := u.subModuleRepo.FindOneBy(map[string]interface{}{
		"id": subject.SubModuleID,
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

	result := &response.MaterialResponse{
		UUID:       subject.UUID,
		Module:     module.Name,
		SubModule:  subModule.Name,
		Subject:    subject.Name,
		SubSubject: make([]*response.SubSubjectResponse, len(subSubjects)),
	}

	for i, subSubject := range subSubjects {
		result.SubSubject[i] = &response.SubSubjectResponse{
			UUID:    subSubject.UUID,
			Name:    subSubject.Name,
			Content: subSubject.Content,
		}
	}

	return result, nil
}
