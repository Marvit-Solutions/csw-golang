package module

import (
	"csw-golang/internal/domain/entity/dto"
)

// ListModule ...
func (mr *moduleRepo) ListModule() ([]dto.ModuleResponse, error) {
	var moduleList []dto.ModuleResponse
	err := mr.db.Preload("SubModule").Preload("Subject").Preload("Exercise").Find(&moduleList).Error
	if err != nil {
		return nil, err
	}

	var moduleResponses []dto.ModuleResponse
	for _, module := range moduleList {
		moduleResponse := dto.ModuleResponse{
			ID:          module.ID,
			Name:        module.Name,
			Description: module.Description,
		}
		for _, subject := range module.Subject {
			subjectResponse := struct {
				ID   string `json:"ID" form:"ID"`
				Name string `json:"Name" form:"Name"`
			}{
				ID:   subject.ID,
				Name: subject.Name,
			}
			moduleResponse.Subject = append(moduleResponse.Subject, subjectResponse)
		}
		for _, excercise := range module.Exercise {
			excerciseResponse := struct {
				ID   string `json:"ID" form:"ID"`
				Name string `json:"Name" form:"Name"`
			}{
				ID:   excercise.ID,
				Name: excercise.Name,
			}
			moduleResponse.Exercise = append(moduleResponse.Exercise, excerciseResponse)
		}

		moduleResponses = append(moduleResponses, moduleResponse)

	}

	return moduleList, nil
}
