package quiz

import (
	"fmt"
	"sort"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) QuizContent(req request.ParamQuizContent) (*response.QuizContentResponse, error) {
	quiz, err := u.quizRepo.FindOneBy(map[string]interface{}{
		"uuid": req.QuizUUID,
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

	quizQuestions, err := u.quizQuestionRepo.FindBy(map[string]interface{}{
		"quiz_id": quiz.ID,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find quizQuestions: %v", err)
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

	mapChoices := make(map[int][]response.OptionItem)

	for _, quizChoice := range quizChoices {
		quizChoiceResponse := &response.OptionItem{
			ID:     quizChoice.ID,
			UUID:   quizChoice.UUID,
			Text:   quizChoice.Content,
			Letter: "",
		}
		mapChoices[quizChoice.QuestionID] = append(mapChoices[quizChoice.QuestionID], *quizChoiceResponse)
	}

	// adding letter
	for _, mapChoice := range mapChoices {
		for i := range mapChoice {
			mapChoice[i].Letter = string(rune('a' + i))
		}
	}

	questionResponse := []response.QuestionItem{}
	k := len(quizQuestions)
	for _, quizQuestion := range quizQuestions {
		questionResponse = append(questionResponse, response.QuestionItem{
			ID:       quizQuestion.ID,
			UUID:     quizQuestion.UUID,
			Question: quizQuestion.Content,
			NoSoal:   k,
			Status:   "belum-dijawab",
			Mark:     quizQuestion.Score,
			Options:  mapChoices[quizQuestion.ID],
		})
		k--
	}

	result := &response.QuizContentResponse{
		ID:             quiz.ID,
		Topic:          quiz.Title,
		Modul:          subject.Name,
		Description:    quiz.Description,
		TotalQuestions: len(quizQuestionsMap),
		TotalTime:      quiz.Time,
		Questions:      questionResponse,
	}

	sort.Slice(result.Questions, func(i, j int) bool {
		return result.Questions[i].ID < result.Questions[j].ID
	})

	return result, nil
}
