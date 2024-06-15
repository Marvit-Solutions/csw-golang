package quizAdmin

import (
	"fmt"
	"time"

	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) Update(req request.QuizAdminPayloadUpdate) error {
	quiz, err := u.quizRepo.FindOneBy(map[string]interface{}{
		"uuid": req.UUID,
	})
	if err != nil {
		return fmt.Errorf("failed to find quiz: %v", err)
	}

	fmt.Println(req.Subject, req.TestType)
	tx := u.db.Begin()
	defer tx.Rollback()

	openTime, _ := time.Parse(time.RFC3339, req.Open)
	closeTime, _ := time.Parse(time.RFC3339, req.Close)
	subject, err := u.subjectRepo.FindOneBy(map[string]interface{}{
		"name": req.Subject,
	})
	if err != nil {
		return fmt.Errorf("failed to find subject: %v", err)
	}

	testType, err := u.testTypeRepo.FindOneBy(map[string]interface{}{
		"name": req.TestType,
	})
	if err != nil {
		return fmt.Errorf("failed to find subject: %v", err)
	}

	// openTime := req.Open
	// closeTime := req.Close

	fmt.Println("tes")
	quizData := &model.Quiz{
		ID:          quiz.ID,
		UUID:        quiz.UUID,
		SubjectID:   subject.ID,
		TestTypeID:  testType.ID,
		Title:       req.Title,
		Description: req.Description,
		Open:        openTime,
		Close:       closeTime,
		Time:        req.Time,
		Attempt:     req.Attempt,
	}
	err = u.quizRepo.Update(quizData, tx)
	if err != nil {
		return fmt.Errorf("failed to Update quiz: %v", err)
	}

	fmt.Println("tes")

	for _, ques := range req.Questions {

		quizQuestion, err := u.quizQuestionRepo.FindOneBy(map[string]interface{}{
			"uuid": ques.UUID,
		})
		if err != nil {
			return fmt.Errorf("failed to find quiz: %v", err)
		}

		questionData := &model.QuizQuestion{
			ID:          quizQuestion.ID,
			UUID:        ques.UUID,
			QuizID:      quiz.ID,
			Content:     ques.Content,
			Score:       ques.Score,
			Explanation: ques.Explanation,
		}

		err = u.quizQuestionRepo.Update(questionData, tx)
		if err != nil {
			return fmt.Errorf("failed to Update quiz question: %v", err)
		}

		for _, ch := range ques.Choices {
			quizChoice, err := u.quizChoiceRepo.FindOneBy(map[string]interface{}{
				"uuid": ch.UUID,
			})
			if err != nil {
				return fmt.Errorf("failed to find quiz: %v", err)
			}
			choice := &model.QuizChoice{
				ID:         quizChoice.ID,
				UUID:       ch.UUID,
				QuestionID: quizQuestion.ID,
				Content:    ch.Content,
				Score:      ch.Score,
				IsCorrect:  ch.IsCorrect,
			}
			err = u.quizChoiceRepo.Update(choice, tx)
			if err != nil {
				return fmt.Errorf("failed to Update quiz choices: %v", err)
			}

		}

	}

	fmt.Println("tes")
	tx.Commit()

	return nil
}
