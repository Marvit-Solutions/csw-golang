package localservice

import (
	"fmt"
	"sort"

	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/repository"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/student_module/domain/localrepository"
	"gorm.io/gorm"
)

type Module struct {
	DB             *gorm.DB
	subjectRepo    repository.SubjectRepository
	subSubjectRepo repository.SubSubjectRepository
	subModuleRepo  repository.SubModuleRepository
	moduleRepo     repository.ModuleRepository
}

func NewModuleService(
	DB *gorm.DB,
	subjectRepo repository.SubjectRepository,
	subSubjectRepo repository.SubSubjectRepository,
	subModuleRepo repository.SubModuleRepository,
	moduleRepo repository.ModuleRepository,
) localrepository.Module {
	return &Module{
		DB,
		subjectRepo,
		subSubjectRepo,
		subModuleRepo,
		moduleRepo,
	}
}

func (svc *Module) FindSubjectInfo(result *response.MaterialFindResponse) ([]*model.Subject, []*model.SubModule, []*model.Module, map[int]int, error) {
	subjects, err := svc.subjectRepo.FindBy(map[string]interface{}{}, 0, 0)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to find subjects: %v", err)
	}

	subjectIDs := make([]int, len(subjects))
	for i, subject := range subjects {
		subjectIDs[i] = subject.ID
	}

	subModuleIDs := make([]int, len(subjects))
	for i, subject := range subjects {
		subModuleIDs[i] = subject.SubModuleID
	}

	subModules, err := svc.subModuleRepo.FindBy(map[string]interface{}{
		"id": subModuleIDs,
	}, 0, 0)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to find sub modules: %v", err)
	}

	subModuleMap := make(map[int]*model.SubModule)
	for _, subModule := range subModules {
		subModuleMap[subModule.ID] = subModule
	}

	moduleIDs := make([]int, len(subjects))
	for i, subModule := range subModules {
		moduleIDs[i] = subModule.ModuleID
	}

	modules, err := svc.moduleRepo.FindBy(map[string]interface{}{
		"id": moduleIDs,
	}, 0, 0)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to find modules: %v", err)
	}

	moduleMap := make(map[int]*model.Module)
	for _, module := range modules {
		moduleMap[module.ID] = module
	}

	countBySubjectID := make(map[int]int)
	for _, subjectID := range subjectIDs {
		count := svc.subSubjectRepo.Count(map[string]interface{}{
			"subject_id": subjectID,
		})

		countBySubjectID[subjectID] = count
	}

	return subjects, subModules, modules, countBySubjectID, nil
}

func (svc *Module) MapSubjectInfo(subjects []*model.Subject, subModules []*model.SubModule, modules []*model.Module, countBySubjectID map[int]int) ([]*response.SubjectFindResponse, error) {
	subModuleMap := make(map[int]*model.SubModule)
	for _, subModule := range subModules {
		subModuleMap[subModule.ID] = subModule
	}

	moduleMap := make(map[int]*model.Module)
	for _, module := range modules {
		moduleMap[module.ID] = module
	}

	result := make([]*response.SubjectFindResponse, len(subjects))
	for i, subject := range subjects {
		subModule, subModuleExists := subModuleMap[subject.SubModuleID]
		if !subModuleExists {
			return nil, fmt.Errorf("failed to find submodule with ID: %d", subject.SubModuleID)
		}

		module, moduleExists := moduleMap[subModule.ModuleID]
		if !moduleExists {
			return nil, fmt.Errorf("failed to find module with ID: %d", subModule.ModuleID)
		}

		result[i] = &response.SubjectFindResponse{
			ID:            subject.ID,
			UUID:          subject.UUID,
			Name:          subject.Name,
			Module:        module.Name,
			SubModule:     subModule.Name,
			CountMaterial: countBySubjectID[subject.ID],
			LastUpdated:   helper.FormatTimeToStringPtr(subject.UpdatedAt),
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result, nil
}

func (svc *Module) FindSubSubjectInfo(result *response.MaterialFindResponse) ([]*model.Subject, []*model.SubSubject, []*model.SubModule, []*model.Module, error) {
	subSubjects, err := svc.subSubjectRepo.FindBy(map[string]interface{}{}, 0, 0)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to find sub subjects: %v", err)
	}

	subjectIDs := make([]int, len(subSubjects))
	for i, subSubject := range subSubjects {
		subjectIDs[i] = subSubject.SubjectID
	}

	subjects, err := svc.subjectRepo.FindBy(map[string]interface{}{
		"id": subjectIDs,
	}, 0, 0)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to find subjects: %v", err)
	}

	subjectMap := make(map[int]*model.Subject)
	for _, subject := range subjects {
		subjectMap[subject.ID] = subject
	}

	subModuleIDs := make([]int, len(subjects))
	for i, subject := range subjects {
		subModuleIDs[i] = subject.SubModuleID
	}

	subModules, err := svc.subModuleRepo.FindBy(map[string]interface{}{
		"id": subModuleIDs,
	}, 0, 0)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to find sub modules: %v", err)
	}

	moduleIDs := make([]int, len(subjects))
	for i, subModule := range subModules {
		moduleIDs[i] = subModule.ModuleID
	}

	modules, err := svc.moduleRepo.FindBy(map[string]interface{}{
		"id": moduleIDs,
	}, 0, 0)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to find modules: %v", err)
	}

	return subjects, subSubjects, subModules, modules, nil
}

func (svc *Module) MapSubSubjectInfo(subjects []*model.Subject, subSubjects []*model.SubSubject, subModules []*model.SubModule, modules []*model.Module) ([]*response.SubSubjectFindResponse, error) {
	subjectMap := make(map[int]*model.Subject)
	for _, subject := range subjects {
		subjectMap[subject.ID] = subject
	}

	subModuleMap := make(map[int]*model.SubModule)
	for _, subModule := range subModules {
		subModuleMap[subModule.ID] = subModule
	}

	moduleMap := make(map[int]*model.Module)
	for _, module := range modules {
		moduleMap[module.ID] = module
	}

	result := make([]*response.SubSubjectFindResponse, len(subSubjects))
	for i, subSubject := range subSubjects {
		subject := subjectMap[subSubject.SubjectID]
		if subject == nil {
			return nil, fmt.Errorf("failed to find subject with ID: %d", subSubject.SubjectID)
		}

		subModule := subModuleMap[subject.SubModuleID]
		if subModule == nil {
			return nil, fmt.Errorf("failed to find submodule with ID: %d", subject.SubModuleID)
		}

		result[i] = &response.SubSubjectFindResponse{
			ID:          subSubject.ID,
			UUID:        subSubject.UUID,
			Name:        subSubject.Name,
			Module:      moduleMap[subModule.ModuleID].Name,
			SubModule:   subModuleMap[subject.SubModuleID].Name,
			Subject:     subjectMap[subSubject.SubjectID].Name,
			LastUpdated: helper.FormatTimeToStringPtr(subSubject.UpdatedAt),
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result, nil
}
