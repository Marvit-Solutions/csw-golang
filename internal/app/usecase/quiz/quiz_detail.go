package quiz

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
)

func (u *usecase) QuizDetail(req request.ParamQuizDetail) (*response.QuizDetailResponse, error) {
	quiz, err := u.quizRepo.FindOneBy(map[string]interface{}{
		"uuid": req.QuizUUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find subject: %v", err)
	}

	quizQuestionTotal := u.quizQuestionRepo.Count(map[string]interface{}{
		"quiz_id": quiz.ID,
	})
	fmt.Println("ini quizQuestionTotal")
	fmt.Println(quizQuestionTotal)

	// find subject name
	// tmpQuizSubjectID := "ce75821b-8641-4705-896c-5175c0fa9ce0"
	subject, err := u.subjectRepo.FindOneBy(map[string]interface{}{
		"id": quiz.SubjectID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find subject: %v", err)
	}

	quizSubmission, err := u.quizSubmissionRepo.FindBy(map[string]interface{}{
		"quiz_id": quiz.ID,
		"user_id": 40,
	}, 0, 0)

	if err != nil {
		return nil, fmt.Errorf("failed to find subject: %v", err)
	}

	// cek sudah dikerjakan atau belum
	var status response.QuizStatus
	// quizSubmissionCount := u.quizSubmissionRepo.Count(map[string]interface{}{
	// 	"quiz_id": quiz.ID,
	// 	"user_id": req.UserID,
	// })
	quizSubmissionCount := len(quizSubmission)
	quizSubmissionUUID := ""
	fmt.Println("ini quizSubmissionCount")
	fmt.Println(quizSubmissionCount)
	if quizSubmissionCount > 0 {
		status = response.SudahDikerjakan
		quizSubmissionUUID = quizSubmission[quizSubmissionCount-1].UUID
	} else {
		status = response.BelumDikerjakan
	}
	fmt.Println("*quiz.Attempt")

	fmt.Println(*quiz.Attempt)
	fmt.Println(quiz.UUID)

	result := &response.QuizDetailResponse{
		ID:                 quiz.ID,
		UUID:               quiz.UUID,
		Subject:            subject.Name,
		Modul:              quiz.Title,
		Description:        quiz.Description,
		TotalQuestions:     quizQuestionTotal,
		TotalTime:          quiz.Time,
		Status:             status,
		AttemptAllowed:     *quiz.Attempt,
		QuizSubmissionUUID: quizSubmissionUUID,
		Score:              5,
		ScoreMax:           *quiz.MaxScore,
		Attempt:            quizSubmissionCount,
	}
	fmt.Println("ini result")

	fmt.Println(result)

	return result, nil

}
