package module

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
)

func (u *usecase) MaterialFind(req request.Material) (*response.MaterialFindResponse, error) {
	result := &response.MaterialFindResponse{}

	var err error
	if *req.Subject && *req.SubSubject || !*req.Subject && !*req.SubSubject {
		result.Subjects, err = u.findAndMapSubjectInfo(result)
		if err != nil {
			return nil, fmt.Errorf("failed to find and map Subject Info: %v", err)
		}

		result.SubSubjects, err = u.findAndMapSubSubjectInfo(result)
		if err != nil {
			return nil, fmt.Errorf("failed to find and map Sub Subject Info: %v", err)
		}
	} else if *req.Subject {
		result.Subjects, err = u.findAndMapSubjectInfo(result)
		if err != nil {
			return nil, fmt.Errorf("failed to find and map Subject Info: %v", err)
		}
	} else if *req.SubSubject {
		result.SubSubjects, err = u.findAndMapSubSubjectInfo(result)
		if err != nil {
			return nil, fmt.Errorf("failed to find and map Sub Subject Info: %v", err)
		}
	}

	return result, nil
}

func (u *usecase) findAndMapSubjectInfo(result *response.MaterialFindResponse) ([]*response.SubjectFindResponse, error) {
	subjects, subModules, modules, countBySubjectID, err := u.moduleLocalRepo.FindSubjectInfo(result)
	if err != nil {
		return nil, fmt.Errorf("failed to find Subject Info: %v", err)
	}

	return u.moduleLocalRepo.MapSubjectInfo(subjects, subModules, modules, countBySubjectID)
}

func (u *usecase) findAndMapSubSubjectInfo(result *response.MaterialFindResponse) ([]*response.SubSubjectFindResponse, error) {
	subjects, subSubjects, subModules, modules, err := u.moduleLocalRepo.FindSubSubjectInfo(result)
	if err != nil {
		return nil, fmt.Errorf("failed to find Sub Subject Info: %v", err)
	}

	return u.moduleLocalRepo.MapSubSubjectInfo(subjects, subSubjects, subModules, modules)
}
