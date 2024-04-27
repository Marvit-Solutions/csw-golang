package module

import (
	"fmt"
	"sort"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) MaterialFind(req request.Material) (*response.MaterialFindResponse, error) {
	result := &response.MaterialFindResponse{}

	if req.Subject && req.SubSubject || !req.Subject && !req.SubSubject {
		subjects, err := u.subjectRepo.FindBy(map[string]interface{}{}, 0, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to find subjects: %v", err)
		}
		subSubjects, err := u.subSubjectRepo.FindBy(map[string]interface{}{}, 0, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to find sub subjects: %v", err)
		}

		subjectIDs := make([]int, len(subjects))
		for i, subject := range subjects {
			subjectIDs[i] = subject.ID
		}

		subModuleIDs := make([]int, len(subjects))
		for i, subject := range subjects {
			subModuleIDs[i] = subject.SubModuleID
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

		moduleIDs := make([]int, len(subjects))
		for i, subModule := range subModules {
			moduleIDs[i] = subModule.ModuleID
		}

		modules, err := u.moduleRepo.FindBy(map[string]interface{}{
			"id": moduleIDs,
		}, 0, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to find modules: %v", err)
		}

		moduleMap := make(map[int]*model.Module)
		for _, module := range modules {
			moduleMap[module.ID] = module
		}

		countBySubjectID := make(map[int]int)
		for _, subjectID := range subjectIDs {
			count := u.subSubjectRepo.Count(map[string]interface{}{
				"subject_id": subjectID,
			})

			countBySubjectID[subjectID] = count
		}

		result.Subjects = make([]*response.SubjectFindResponse, len(subjects))
		for i, subject := range subjects {
			subModule := subModuleMap[subject.SubModuleID]
			if subModule == nil {
				return nil, fmt.Errorf("failed to find submodule with ID: %d", subject.SubModuleID)
			}

			module := moduleMap[subModule.ModuleID]
			if module == nil {
				return nil, fmt.Errorf("failed to find module with ID: %d", subModule.ModuleID)
			}

			result.Subjects[i] = &response.SubjectFindResponse{
				ID:            subject.ID,
				UUID:          subject.UUID,
				Name:          subject.Name,
				Module:        module.Name,
				SubModule:     subModuleMap[subject.SubModuleID].Name,
				CountMaterial: countBySubjectID[subject.ID],
				LastUpdated:   helper.FormatTimeToStringPtr(subject.UpdatedAt),
			}
		}

		subjectIDs = make([]int, len(subSubjects))
		for i, subSubject := range subSubjects {
			subjectIDs[i] = subSubject.SubjectID
		}

		subjects, err = u.subjectRepo.FindBy(map[string]interface{}{
			"id": subjectIDs,
		}, 0, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to find subjects: %v", err)
		}

		subjectMap := make(map[int]*model.Subject)
		for _, subject := range subjects {
			subjectMap[subject.ID] = subject
		}

		result.SubSubjects = make([]*response.SubSubjectFindResponse, len(subSubjects))
		for i, subSubject := range subSubjects {
			subject := subjectMap[subSubject.SubjectID]
			if subject == nil {
				return nil, fmt.Errorf("failed to find subject with ID: %d", subSubject.SubjectID)
			}

			subModule := subModuleMap[subject.SubModuleID]
			if subModule == nil {
				return nil, fmt.Errorf("failed to find submodule with ID: %d", subject.SubModuleID)
			}

			result.SubSubjects[i] = &response.SubSubjectFindResponse{
				ID:          subSubject.ID,
				UUID:        subSubject.UUID,
				Name:        subSubject.Name,
				Module:      moduleMap[subModule.ModuleID].Name,
				SubModule:   subModuleMap[subject.SubModuleID].Name,
				Subject:     subjectMap[subSubject.SubjectID].Name,
				LastUpdated: helper.FormatTimeToStringPtr(subSubject.UpdatedAt),
			}
		}

		sort.Slice(result.Subjects, func(i, j int) bool {
			return result.Subjects[i].ID < result.Subjects[j].ID
		})

		sort.Slice(result.SubSubjects, func(i, j int) bool {
			return result.SubSubjects[i].ID < result.SubSubjects[j].ID
		})

	} else if req.Subject {
		subjects, err := u.subjectRepo.FindBy(map[string]interface{}{}, 0, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to find subjects: %v", err)
		}

		subModuleIDs := make([]int, len(subjects))
		for i, subject := range subjects {
			subModuleIDs[i] = subject.SubModuleID
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

		moduleIDs := make([]int, len(subjects))
		for i, subModule := range subModules {
			moduleIDs[i] = subModule.ModuleID
		}

		modules, err := u.moduleRepo.FindBy(map[string]interface{}{
			"id": moduleIDs,
		}, 0, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to find modules: %v", err)
		}

		moduleMap := make(map[int]*model.Module)
		for _, module := range modules {
			moduleMap[module.ID] = module
		}

		subjectIDs := make([]int, len(subjects))
		for i, subject := range subjects {
			subjectIDs[i] = subject.ID
		}

		countBySubjectID := make(map[int]int)
		for _, subjectID := range subjectIDs {
			count := u.subSubjectRepo.Count(map[string]interface{}{
				"subject_id": subjectID,
			})

			countBySubjectID[subjectID] = count
		}

		result.Subjects = make([]*response.SubjectFindResponse, len(subjects))
		for i, subject := range subjects {
			subModule := subModuleMap[subject.SubModuleID]
			if subModule == nil {
				return nil, fmt.Errorf("failed to find submodule with ID: %d", subject.SubModuleID)
			}

			module := moduleMap[subModule.ModuleID]
			if module == nil {
				return nil, fmt.Errorf("failed to find module with ID: %d", subModule.ModuleID)
			}

			result.Subjects[i] = &response.SubjectFindResponse{
				ID:            subject.ID,
				UUID:          subject.UUID,
				Name:          subject.Name,
				Module:        module.Name,
				SubModule:     subModuleMap[subject.SubModuleID].Name,
				CountMaterial: countBySubjectID[subject.ID],
				LastUpdated:   helper.FormatTimeToStringPtr(subject.UpdatedAt),
			}
		}

		sort.Slice(result.Subjects, func(i, j int) bool {
			return result.Subjects[i].ID < result.Subjects[j].ID
		})

	} else if req.SubSubject {
		subSubjects, err := u.subSubjectRepo.FindBy(map[string]interface{}{}, 0, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to find sub subjects: %v", err)
		}

		subjectIDs := make([]int, len(subSubjects))
		for i, subSubject := range subSubjects {
			subjectIDs[i] = subSubject.SubjectID
		}

		subjects, err := u.subjectRepo.FindBy(map[string]interface{}{
			"id": subjectIDs,
		}, 0, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to find subjects: %v", err)
		}

		subjectMap := make(map[int]*model.Subject)
		for _, subject := range subjects {
			subjectMap[subject.ID] = subject
		}

		subModuleIDs := make([]int, len(subjects))
		for i, subject := range subjects {
			subModuleIDs[i] = subject.SubModuleID
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

		moduleIDs := make([]int, len(subjects))
		for i, subModule := range subModules {
			moduleIDs[i] = subModule.ModuleID
		}

		modules, err := u.moduleRepo.FindBy(map[string]interface{}{
			"id": moduleIDs,
		}, 0, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to find modules: %v", err)
		}

		moduleMap := make(map[int]*model.Module)
		for _, module := range modules {
			moduleMap[module.ID] = module
		}

		result.SubSubjects = make([]*response.SubSubjectFindResponse, len(subSubjects))
		for i, subSubject := range subSubjects {
			subject := subjectMap[subSubject.SubjectID]
			if subject == nil {
				return nil, fmt.Errorf("failed to find subject with ID: %d", subSubject.SubjectID)
			}

			subModule := subModuleMap[subject.SubModuleID]
			if subModule == nil {
				return nil, fmt.Errorf("failed to find submodule with ID: %d", subject.SubModuleID)
			}

			result.SubSubjects[i] = &response.SubSubjectFindResponse{
				ID:          subSubject.ID,
				UUID:        subSubject.UUID,
				Name:        subSubject.Name,
				Module:      moduleMap[subModule.ModuleID].Name,
				SubModule:   subModuleMap[subject.SubModuleID].Name,
				Subject:     subjectMap[subSubject.SubjectID].Name,
				LastUpdated: helper.FormatTimeToStringPtr(subSubject.UpdatedAt),
			}
		}

		sort.Slice(result.SubSubjects, func(i, j int) bool {
			return result.SubSubjects[i].ID < result.SubSubjects[j].ID
		})
	}

	return result, nil
}
