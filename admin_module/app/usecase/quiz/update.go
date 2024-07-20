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

	fmt.Println("tesusecase1")
	openTime, _ := time.Parse(time.RFC3339, req.Open)
	closeTime, _ := time.Parse(time.RFC3339, req.Close)
	subject, err := u.subjectRepo.FindOneBy(map[string]interface{}{
		"name": req.Subject,
	})
	if err != nil {
		return fmt.Errorf("failed to find subject: %v", err)
	}

	fmt.Println("tesusecase2")
	testType, err := u.testTypeRepo.FindOneBy(map[string]interface{}{
		"name": req.TestType,
	})
	if err != nil {
		return fmt.Errorf("failed to find subject: %v", err)
	}

	// openTime := req.Open
	// closeTime := req.Close

	fmt.Println("tesusecase3")
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

	fmt.Println("tesusecase4")

	for _, ques := range req.Questions {
		fmt.Println("tesusecase5")

		if ques.UUID == "" {
			fmt.Println("tesusecase55")
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

			fmt.Println("tesusecase555")

		} else {
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
			fmt.Println("tesusecase6")

			err = u.quizQuestionRepo.Update(questionData, tx)
			if err != nil {
				return fmt.Errorf("failed to Update quiz question: %v", err)
			}
			fmt.Println("tesusecase7")

			for _, ch := range ques.Choices {
				if ch.UUID == "" {
					fmt.Println("tesusecase77")

					choice := &model.QuizChoice{
						QuestionID: questionData.ID,
						Content:    ch.Content,
						Score:      ch.Score,
						IsCorrect:  ch.IsCorrect,
					}
					_, err := u.quizChoiceRepo.Create(choice, tx)
					if err != nil {
						return fmt.Errorf("failed to create quiz choices: %v", err)
					}

				} else {
					quizChoice, err := u.quizChoiceRepo.FindOneBy(map[string]interface{}{
						"uuid": ch.UUID,
					})
					if err != nil {
						return fmt.Errorf("failed to find quiz: %v", err)
					}
					fmt.Println(ch.Content)
					fmt.Println("tesusecase8")

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
					fmt.Println("tesusecase9")
				}

			}
		}

	}

	fmt.Println("tes update")
	tx.Commit()

	return nil
}
