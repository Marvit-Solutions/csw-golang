package pretest

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/helper/time"
	"fmt"
)

func (pr *pretestUsecase) GetAllPretests(userID, module string) (error, []dto.GetAllPretestResponse) {
	var allPretests []dto.GetAllPretestResponse

	err, pretests := pr.pretestRepo.GetAllPretests(userID, module)
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
						MeetingDate: time.ConvertTimeFormat(SubjectTestTypeQuiz.MeetingDate),
						Open:        time.ConvertTimeFormat(SubjectTestTypeQuiz.Open),
						Description: SubjectTestTypeQuiz.Description,
						Time:        SubjectTestTypeQuiz.Time,
						Point:       SubjectTestTypeQuiz.Point,
						// Status:      SubjectTestTypeQuiz.Status,
						Attempt: SubjectTestTypeQuiz.Attempt,
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

func (pr *pretestUsecase) GetPretestById(pretestId string) (error, dto.Pretest) {
	var pretest dto.Pretest

	err, pretests := pr.pretestRepo.GetPretestById(pretestId)
	if err != nil {
		return err, dto.Pretest{}
	}

	pretest = dto.Pretest{
		IDPretest:   pretests.ID,
		TestType:    pretests.TestType,
		Title:       pretests.Title,
		MeetingDate: time.ConvertTimeFormat(pretests.MeetingDate),
		Open:        time.ConvertTimeFormat(pretests.Open),
		Description: pretests.Description,
		Time:        pretests.Time,
		Point:       pretests.Point,
		// Status:      pretests.Status,
		Attempt: pretests.Attempt,
	}

	for _, question := range pretests.QuestionQuizzes {
		questionResponse := dto.QuestionPretest{
			IDQuestion: question.ID,
			Image:      question.Image,
			Content:    question.Content,
			Weight:     question.Weight,
			// Status:     question.Status,
		}

		for _, choice := range question.ChoiceQuizzes {
			choiceResponse := dto.ChoicePretest{
				IDChoice:  choice.ID,
				Content:   choice.Content,
				IsCorrect: choice.IsCorrect,
				Weight:    choice.Weight,
			}

			questionResponse.Choice = append(questionResponse.Choice, &choiceResponse)
		}

		pretest.Question = append(pretest.Question, &questionResponse)
	}

	return nil, pretest
}
func (pr *pretestUsecase) GetPretestReview(pretestId, status string) (error, dto.PretestReviewResponse) {
	var pretest dto.PretestReviewResponse

	err, pretests := pr.pretestRepo.GetPretestReview(pretestId, status)
	if err != nil {
		return err, dto.PretestReviewResponse{}
	}

	// belum ada handle jawaban tiap user
	pretest = dto.PretestReviewResponse{
		IDPretest:      pretests.ID,
		TestType:       pretests.TestType,
		Title:          pretests.Title,
		StartTime:      time.ConvertTimeFormat(pretests.Open),
		SubmissionTime: time.ConvertTimeFormat(pretests.GradeQuizzes.GradingTime),
		Description:    pretests.Description,
		Time:           pretests.Time,
		Point:          fmt.Sprintf("%d dari %d", pretests.GradeQuizzes.Score, pretests.Point),
		// Status:         pretests.Status,
		Attempt: pretests.Attempt,
	}

	for _, question := range pretests.QuestionQuizzes {
		questionResponse := dto.QuestionPretest{
			IDQuestion: question.ID,
			Image:      question.Image,
			Content:    question.Content,
			Weight:     question.Weight,
			// Status:     question.Status,
		}

		for _, choice := range question.ChoiceQuizzes {
			choiceResponse := dto.ChoicePretest{
				IDChoice:  choice.ID,
				Content:   choice.Content,
				IsCorrect: choice.IsCorrect,
				Weight:    choice.Weight,
			}

			for _, userSubmittedAnswer := range choice.UserSubmittedAnswerQuizzes {
				userSubmittedAnswerResponse := dto.UserSubmittedAnswerPretest{
					IDUserSubmittedAnswerPretest: userSubmittedAnswer.ID,
				}

				choiceResponse.UserSubmittedAnswer = append(choiceResponse.UserSubmittedAnswer, &userSubmittedAnswerResponse)
			}

			if len(choiceResponse.UserSubmittedAnswer) > 0 {
				choiceResponse.Selected = true
			} else {
				choiceResponse.Selected = false
			}

			questionResponse.Choice = append(questionResponse.Choice, &choiceResponse)
		}

		pretest.Question = append(pretest.Question, &questionResponse)
	}

	return nil, pretest
}

func (pr *pretestUsecase) SubmitPretest(id string, req dto.PretestSubmitRequest) error {
	err := pr.pretestRepo.SubmitPretest(id, req)
	if err != nil {
		return err
	}

	return nil
}
func (pr *pretestUsecase) GradingPretest(id string, req dto.PretestSubmitRequest) error {
	err := pr.pretestRepo.GradingPretest(id, req)
	if err != nil {
		return err
	}

	return nil
}
