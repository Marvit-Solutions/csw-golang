package quizAdmin

import (
	"fmt"
	"sort"

	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) QuizAdminUpdateDetail(req request.ParamQuizAdminUpdateDetail) (*response.QuizAdminUpdateDetailResponse, error) {
	quiz, err := u.quizRepo.FindOneBy(map[string]interface{}{
		"uuid": req.UUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find quiz: %v", err)
	}

	subject, err := u.subjectRepo.FindOneBy(map[string]interface{}{
		"id": quiz.SubjectID,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to find subject: %v", err)
	}

	testType, err := u.testTypeRepo.FindOneBy(map[string]interface{}{
		"id": quiz.TestTypeID,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to find subject: %v", err)
	}

	quizQuestions, err := u.quizQuestionRepo.FindBy(map[string]interface{}{
		"quiz_id": quiz.ID,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find quiz questions: %v", err)
	}

	quizQuestionsMap := make(map[int]*model.QuizQuestion)
	for _, quizQuestion := range quizQuestions {
		quizQuestionsMap[quizQuestion.ID] = quizQuestion
	}

	quizQuestionIDs := make([]int, len(quizQuestions))
	for i, quizQuestion := range quizQuestions {
		quizQuestionIDs[i] = quizQuestion.ID
	}

	// get all choices
	quizChoices, err := u.quizChoiceRepo.FindBy(map[string]interface{}{
		"question_id": quizQuestionIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find quiz choices: %v", err)
	}

	mapChoices := make(map[int][]response.ChoiceQuizAdmin)

	for _, quizChoice := range quizChoices {
		quizChoiceResponse := &response.ChoiceQuizAdmin{
			UUID:      quizChoice.UUID,
			Content:   quizChoice.Content,
			Score:     quizChoice.Score,
			IsCorrect: quizChoice.IsCorrect,
		}
		mapChoices[quizChoice.QuestionID] = append(mapChoices[quizChoice.QuestionID], *quizChoiceResponse)
	}

	questionResponse := []response.QuestionQuizAdmin{}
	k := len(quizQuestions)
	for _, quizQuestion := range quizQuestions {
		questionResponse = append(questionResponse, response.QuestionQuizAdmin{
			UUID:        quizQuestion.UUID,
			Content:     quizQuestion.Content,
			Score:       quizQuestion.Score,
			Explanation: quizQuestion.Explanation,
			Choices:     mapChoices[quizQuestion.ID],
		})
		k--
	}

	result := &response.QuizAdminUpdateDetailResponse{
		UUID:        req.UUID,
		Title:       quiz.Title,
		Subject:     subject.Name,
		TestType:    testType.Name,
		Time:        quiz.Time,
		Attempt:     quiz.Attempt,
		Open:        quiz.Open,
		Close:       quiz.Close,
		Description: quiz.Description,
		Questions:   questionResponse,
	}

	sort.Slice(result.Questions, func(i, j int) bool {
		return result.Questions[i].ID < result.Questions[j].ID
	})

	return result, nil
}
