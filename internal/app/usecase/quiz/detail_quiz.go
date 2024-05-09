package quiz

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
)

func (u *usecase) DetailQuiz(req request.ParamDetailQuiz) (*response.DetailQuizResponse, error) {
	quiz, err := u.quizRepo.FindOneBy(map[string]interface{}{
		"uuid": req.QuizUUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find subject: %v", err)
	}

	quizQuestionTotal := u.quizQuestionRepo.Count(map[string]interface{}{
		"quiz_id": quiz.ID,
	})
	fmt.Println(quizQuestionTotal)
	if err != nil {
		return nil, fmt.Errorf("failed to find sub subjects: %v", err)
	}

	// find subject name
	subject, err := u.subjectRepo.FindOneBy(map[string]interface{}{
		"uuid": quiz.S,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find subject: %v", err)
	}

	result := &response.DetailQuizResponse{
		ID:             quiz.ID,
		Subject:        subject.,
		Modul:          quiz.Title,
		TotalQuestions: quizQuestionTotal,
		TotalTime:      quiz.Time,
		Questions:      questionResponse,
	}

	return result, nil

}
