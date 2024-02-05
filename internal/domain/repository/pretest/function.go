package pretest

import (
	"csw-golang/internal/domain/entity/datastruct"
	"csw-golang/internal/domain/entity/dto"
)

func (pr pretestRepo) GetAllPretests() (error, []dto.GetAllPretestResponse) {
	var allPretests []dto.GetAllPretestResponse
	var modules []datastruct.Modules

	err := pr.db.Preload("SubModules").Preload("SubModules.Subjects").Preload("SubModules.Subjects.SubjectTestTypeQuizzes").Find(&modules).Error
	if err != nil {
		return err, allPretests
	}

	for _, module := range modules {
		allPretest := dto.GetAllPretestResponse{
			ID:   module.ID,
			Name: module.Name,
		}

		for _, subModule := range module.SubModules {
			subModuleResponse := dto.SubModulePretest{
				IDSubModule:          subModule.ID,
				NameSubModule:        subModule.Name,
				DescriptionSubModule: subModule.Description,
			}

			for _, subject := range subModule.Subjects {
				subjectResponse := dto.SubjectPretest{
					IDSubject:   subject.ID,
					NameSubject: subject.Name,
				}

				for _, SubjectTestTypeQuiz := range subject.SubjectTestTypeQuizzes {
					subjectTestTypeQuizResponse := dto.Pretest{
						IDPretest:   SubjectTestTypeQuiz.ID,
						TestType:    SubjectTestTypeQuiz.TestType,
						Title:       SubjectTestTypeQuiz.Title,
						MeetingDate: SubjectTestTypeQuiz.MeetingDate,
						Open:        SubjectTestTypeQuiz.Open,
						Description: SubjectTestTypeQuiz.Description,
						Time:        SubjectTestTypeQuiz.Time,
						Point:       SubjectTestTypeQuiz.Point,
						Status:      SubjectTestTypeQuiz.Status,
						Attempt:     SubjectTestTypeQuiz.Attempt,
					}

					subjectResponse.Pretest = append(subjectResponse.Pretest, &subjectTestTypeQuizResponse)
				}
				subModuleResponse.Subject = append(subModuleResponse.Subject, &subjectResponse)
			}

			allPretest.SubModule = append(allPretest.SubModule, &subModuleResponse)
		}
		allPretests = append(allPretests, allPretest)
	}

	return nil, allPretests
}
