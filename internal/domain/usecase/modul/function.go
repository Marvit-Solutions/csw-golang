package module

import (
	"csw-golang/internal/domain/entity/dto"
	"fmt"
)

func (mod *moduleUsecase) GetListModules() ([]dto.ModuleResponse, error) {
	fmt.Println("GetListModules usecase")
	moduleList, err := mod.moduleRepo.GetListModules()
	if err != nil {
		return nil, err
	}
	quizList, err := mod.moduleRepo.GetListSubjectQuiz()

	var moduleResponses []dto.ModuleResponse
	for _, module := range moduleList {
		moduleResponse := dto.ModuleResponse{
			ID:          module.ID,
			Name:        module.Name,
			Description: module.Description,
		}
		for _, subject := range module.Subjects {
			subjectResponse := struct {
				ID   string `json:"ID" form:"ID"`
				Name string `json:"Name" form:"Name"`
			}{
				ID:   subject.ID,
				Name: subject.Name,
			}
			for _, quiz := range quizList {
				if quiz.SubjectID == subject.ID {
					quizResponse := struct {
						ID   string `json:"ID" form:"ID"`
						Name string `json:"Name" form:"Name"`
					}{
						ID:   quiz.ID,
						Name: subject.Name,
					}

					moduleResponse.Exercise = append(moduleResponse.Exercise, quizResponse)
				}
			}
			moduleResponse.Subject = append(moduleResponse.Subject, subjectResponse)
		}
		moduleResponses = append(moduleResponses, moduleResponse)
	}
	return moduleResponses, nil
}

func (mod *moduleUsecase) GetSubjectsBySubmoduleID(submoduleID string) ([]dto.SubjectResponse, error) {
	fmt.Println("GetSubjectsBySubmoduleID usecase")

	subjects, err := mod.moduleRepo.GetSubjectsBySubmoduleID(submoduleID)
	if err != nil {
		return nil, err
	}

	var subjectResponses []dto.SubjectResponse

	for _, subject := range subjects {
		subjectResponse := dto.SubjectResponse{
			ID:   subject.ID,
			Name: subject.Name,
		}
		for _, subSubject := range subject.SubSubject {
			subSubjectResponse := dto.SubSubject{
				ID:      subSubject.ID,
				Name:    subSubject.Name,
				Content: subSubject.Content,
			}
			subjectResponse.SubSubject = append(subjectResponse.SubSubject, subSubjectResponse)
		}

		subjectResponses = append(subjectResponses, subjectResponse)

	}

	return subjectResponses, nil
}
