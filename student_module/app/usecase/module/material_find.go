package module

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/response"
)

func (u *usecase) MaterialFind() (*response.MaterialFindResponse, error) {
	result := &response.MaterialFindResponse{}

	// get all subject
	subjects, subModules, modules, countBySubjectID, err := u.moduleLocalRepo.FindSubjectInfo(result)
	if err != nil {
		return nil, fmt.Errorf("failed to find Subject Info: %v", err)
	}

	// map subject
	subjectResp, err := u.moduleLocalRepo.MapSubjectInfo(subjects, subModules, modules, countBySubjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to map Subject Info: %v", err)
	}

	// get all sub subject
	subjects, subSubjects, subModules, modules, err := u.moduleLocalRepo.FindSubSubjectInfo(result)
	if err != nil {
		return nil, fmt.Errorf("failed to find Sub Subject Info: %v", err)
	}

	// map sub subject
	subSubjectResp, err := u.moduleLocalRepo.MapSubSubjectInfo(subjects, subSubjects, subModules, modules)
	if err != nil {
		return nil, fmt.Errorf("failed to map Sub Subject Info: %v", err)
	}

	// assign to result
	result.Subjects = subjectResp
	result.SubSubjects = subSubjectResp

	return result, nil
}
