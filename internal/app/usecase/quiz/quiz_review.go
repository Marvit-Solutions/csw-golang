package quiz

import (
	"fmt"
	"sort"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) QuizReview(req request.ParamQuizReview) (*response.QuizReviewResponse, error) {

	quizSubmission, err := u.quizSubmissionRepo.FindOneBy(map[string]interface{}{
		"uuid": req.QuizSubmissionUUID,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to find quiz submission: %v", err)
	}
	quiz, err := u.quizRepo.FindOneBy(map[string]interface{}{
		"id": quizSubmission.QuizID,
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

	quizSubmissionCount := u.quizSubmissionRepo.Count(map[string]interface{}{
		"quiz_id": quiz.ID,
		"user_id": 40,
	})

	quizQuestions, err := u.quizQuestionRepo.FindBy(map[string]interface{}{
		"quiz_id": quiz.ID,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find sub quiz questions: %v", err)
	}

	quizQuestionsMap := make(map[int]*model.QuizQuestion)
	for _, quizQuestion := range quizQuestions {
		quizQuestionsMap[quizQuestion.ID] = quizQuestion
	}

	quizQuestionIDs := make([]int, len(quizQuestions))
	for i, quizQuestion := range quizQuestions {
		quizQuestionIDs[i] = quizQuestion.ID
	}

	quizChoices, err := u.quizChoiceRepo.FindBy(map[string]interface{}{
		"question_id": quizQuestionIDs,
	}, 0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to find quiz choices: %v", err)
	}

	mapChoices := make(map[int][]response.OptionItemReview)

	for _, quizChoice := range quizChoices {
		quizChoiceResponse := &response.OptionItemReview{
			ID:        quizChoice.ID,
			UUID:      quizChoice.UUID,
			Text:      quizChoice.Content,
			Letter:    "",
			IsCorrect: quizChoice.IsCorrect,
		}
		mapChoices[quizChoice.QuestionID] = append(mapChoices[quizChoice.QuestionID], *quizChoiceResponse)
	}

	// adding letter
	for _, mapChoice := range mapChoices {
		for i := range mapChoice {
			mapChoice[i].Letter = string(rune('a' + i))
		}
	}

	userAnswers, err := u.quizLocalRepo.FindUserAnswerReview(quizSubmission.ID)
	userAnswersMap := make(map[int]*response.UserAnswer)
	for _, userAnswer := range userAnswers {
		userAnswersMap[userAnswer.QuestionId] = userAnswer
	}

	if err != nil {
		return nil, fmt.Errorf("failed to find user answer: %v", err)
	}

	questionReviews := make([]response.QuestionsReview, len(quizQuestions))

	k := len(quizQuestions)
	for i, quizQuestion := range quizQuestions {
		// find right answer
		correctOptionIdx := 0
		correctOptionText := ""
		for _, option := range mapChoices[quizQuestion.ID] {
			if option.IsCorrect {
				correctOptionIdx = option.ID
				correctOptionText = option.Text
				break
			}
		}

		userAnswer := 0
		if userAnswersMap[quizQuestion.ID] != nil {
			userAnswer = userAnswersMap[quizQuestion.ID].ChoiceId
		}

		var status response.QuestionReviewItemStatus

		if userAnswer == 0 {
			status = response.BelumDiJawab
		} else {
			status = response.SudahDiJawab
		}

		questionReviews[i] = response.QuestionsReview{
			UserAnswer:      userAnswer,
			RightAnswer:     correctOptionIdx,
			RightAnswerText: correctOptionText,
			QuestionReviewItem: response.QuestionReviewItem{
				ID:          quizQuestion.ID,
				UUID:        quizQuestion.UUID,
				Question:    quizQuestion.Content,
				NoSoal:      k,
				Status:      string(status),
				Mark:        quizQuestion.Score,
				Explanation: quizQuestion.Explanation,
				Options:     mapChoices[quizQuestion.ID],
			},
		}
		k--
	}

	result := &response.QuizReviewResponse{
		ID:                quiz.ID,
		UUID:              req.QuizSubmissionUUID,
		Topic:             quiz.Title,
		Modul:             subject.Name,
		TotalQuestions:    len(quizQuestionsMap),
		StartedAt:         helper.FormatIndonesianDate(quizSubmission.StartedAt),
		TotalTime:         quizSubmission.TimeRequired,
		TotalRightAnswers: quizSubmission.RightAnswer,
		Score:             quizSubmission.Score,
		MaxScore:          quiz.MaxScore,
		Attempt:           quizSubmissionCount,
		Questions:         questionReviews,
	}

	sort.Slice(result.Questions, func(i, j int) bool {
		return result.Questions[i].QuestionReviewItem.ID < result.Questions[j].QuestionReviewItem.ID
	})

	return result, nil
}
