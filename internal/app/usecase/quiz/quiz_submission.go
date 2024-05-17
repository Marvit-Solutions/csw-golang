package quiz

import (
	"fmt"
	"time"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) QuizSubmission(req request.QuizSubmissionRequest) error {

	tx := u.db.Begin()
	defer tx.Rollback()

	quizQuestions, err := u.quizQuestionRepo.FindBy(map[string]interface{}{
		"quiz_id": req.QuizID,
	}, 0, 0)
	if err != nil {
		return fmt.Errorf("failed to find sub subjects: %v", err)
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
		return fmt.Errorf("failed to find sub subjects: %v", err)
	}

	mapChoices := make(map[int][]response.OptionItemSubmission)

	for _, quizChoice := range quizChoices {
		quizChoiceResponse := &response.OptionItemSubmission{
			ID:        quizChoice.ID,
			UUID:      quizChoice.UUID,
			Text:      quizChoice.Content,
			IsCorrect: quizChoice.IsCorrect,
		}
		mapChoices[quizChoice.QuestionID] = append(mapChoices[quizChoice.QuestionID], *quizChoiceResponse)
	}

	totalRightAnswer := 0
	score := 0

	// for _, quizQuestionMap := range quizQuestionsMap {
	// 	if req.Questions[]
	// }

	for _, question := range req.Questions {
		var rightAnswer int
		for _, option := range mapChoices[question.ID] {
			if option.IsCorrect {
				rightAnswer = option.ID
			}
		}
		if question.UserAnswer == rightAnswer {
			totalRightAnswer++
			score = score + question.Mark
		}
	}

	fmt.Println(score)

	quizSubmissionData := &model.QuizSubmission{
		UserID:       req.UserID,
		QuizID:       req.QuizID,
		StartedAt:    time.Now().Add(-time.Hour),
		FinishedAt:   time.Now(),
		TimeRequired: req.TimeRequired,
		Score:        score,
		RightAnswer:  totalRightAnswer,
	}

	quizSubmission, err := u.quizSubmissionRepo.Create(quizSubmissionData, tx)

	if err != nil {
		return fmt.Errorf("failed to create quizSubmission: %v", err)
	}

	// insert user quiz answer and start calculate the score
	for _, ques := range req.Questions {
		// Periksa apakah UserAnswer tidak nil dan tidak sama dengan undefined
		if ques.UserAnswer != 0 {
			quizAnswer := &model.QuizAnswer{
				SubmissionID: quizSubmission.ID,
				ChoiceID:     &ques.UserAnswer, // Menggunakan nilai yang tidak diubah
				IsMarked:     false,
				CreatedBy:    40,
				UpdatedBy:    40,
			}

			quizAnswer, err = u.quizAnswerRepo.Create(quizAnswer, tx)
			if err != nil {
				return fmt.Errorf("failed to create quizAnswer: %v", err)
			}
			fmt.Println(quizAnswer)
		}
	}

	//end off calculate score

	tx.Commit()

	return nil
}
