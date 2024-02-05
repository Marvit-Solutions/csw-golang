package pretest

import (
	"csw-golang/internal/domain/entity/dto"
)

func (pr *pretestUsecase) GetAllPretests() (error, []dto.GetAllPretestResponse) {
	var allPretests []dto.GetAllPretestResponse

	err, pretests := pr.pretestRepo.GetAllPretests()
	if err != nil {
		return err, []dto.GetAllPretestResponse{}
	}

	for _, pretest := range pretests {
		allPretest := dto.GetAllPretestResponse{
			ID:   pretest.ID,
			Name: pretest.Name,
		}

		for _, subModule := range pretest.SubModules {
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
