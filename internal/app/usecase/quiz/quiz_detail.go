package quiz

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
)

func (u *usecase) QuizDetail(req request.ParamQuizDetail) (*response.QuizDetailResponse, error) {
	// get user id

	quiz, err := u.quizRepo.FindOneBy(map[string]interface{}{
		"uuid": req.QuizUUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find subject: %v", err)
	}

	quizQuestionTotal := u.quizQuestionRepo.Count(map[string]interface{}{
		"quiz_id": quiz.ID,
	})

	// find subject name
	// tmpQuizSubjectID := "ce75821b-8641-4705-896c-5175c0fa9ce0"
	subject, err := u.subjectRepo.FindOneBy(map[string]interface{}{
		"id": quiz.SubjectID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find subject: %v", err)
	}

	quizSubmissions, err := u.quizSubmissionRepo.FindBy(map[string]interface{}{
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
	quizSubmissionCount := len(quizSubmissions)
	quizSubmissionUUID := ""
	max_score := 0

	if quizSubmissionCount > 0 {
		status = response.SudahDikerjakan
		// cek jika sudah dikerjakan sekali atau lebih
		if quizSubmissionCount == 1 {
			max_score = quizSubmissions[quizSubmissionCount-1].Score
			quizSubmissionUUID = quizSubmissions[quizSubmissionCount-1].UUID
		} else {
			// cari quizSubmissionUUID dengan nilai yang paling tinggi
			for _, quizSubmission := range quizSubmissions {
				if quizSubmission.Score > max_score {
					max_score = quizSubmission.Score
					quizSubmissionUUID = quizSubmission.UUID
				}
			}
		}

	} else {
		status = response.BelumDikerjakan
	}

	result := &response.QuizDetailResponse{
		ID:                 quiz.ID,
		UUID:               quiz.UUID,
		Subject:            subject.Name,
		Modul:              quiz.Title,
		Description:        quiz.Description,
		TotalQuestions:     quizQuestionTotal,
		TotalTime:          quiz.Time,
		Status:             status,
		AttemptAllowed:     quiz.Attempt,
		QuizSubmissionUUID: quizSubmissionUUID,
		Score:              max_score,
		MaxScore:           quiz.MaxScore,
		Attempt:            quizSubmissionCount,
	}

	return result, nil

}
