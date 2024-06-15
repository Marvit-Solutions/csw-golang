package quizAdmin

import (
	"fmt"
	"time"

	"github.com/Marvit-Solutions/csw-golang/admin_module/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/library/struct/model"
)

func (u *usecase) Create(req request.QuizAdminPayload) error {
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
		SubjectID:   subject.ID,
		TestTypeID:  testType.ID,
		Title:       req.Title,
		Description: req.Description,
		Open:        openTime,
		Close:       closeTime,
		Time:        req.Time,
		Attempt:     req.Attempt,
	}
	quiz, err := u.quizRepo.Create(quizData, tx)
	if err != nil {
		return fmt.Errorf("failed to create quiz: %v", err)
	}

	fmt.Println("tes")

	for _, ques := range req.Questions {
		questionData := &model.QuizQuestion{
			QuizID:      quiz.ID,
			Content:     ques.Content,
			Score:       ques.Score,
			Explanation: ques.Explanation,
		}
		question, err := u.quizQuestionRepo.Create(questionData, tx)
		if err != nil {
			return fmt.Errorf("failed to create quiz question: %v", err)
		}
		for _, ch := range ques.Choices {
			choice := &model.QuizChoice{
				QuestionID: question.ID,
				Content:    ch.Content,
				Score:      ch.Score,
				IsCorrect:  ch.IsCorrect,
			}
			_, err := u.quizChoiceRepo.Create(choice, tx)
			if err != nil {
				return fmt.Errorf("failed to create quiz choices: %v", err)
			}

		}

	}

	fmt.Println("tes")
	tx.Commit()

	return nil
}
