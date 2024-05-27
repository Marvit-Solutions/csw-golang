package quiz

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
)

func (u *usecase) QuizDetail(req request.ParamQuizDetail) (*response.QuizDetailResponse, error) {
	user, err := u.userRepo.FindOneBy(map[string]interface{}{
		"id":      req.AuthenticatedUser,
		"role_id": helper.PembeliPaketBimbel,
	})
	if user == nil {
		return nil, helper.ErrAccessDenied
	}
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %v", err)
	}

	quiz, err := u.quizRepo.FindOneBy(map[string]interface{}{
		"uuid": req.QuizUUID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find quiz: %v", err)
	}

	quizQuestionTotal := u.quizQuestionRepo.Count(map[string]interface{}{
		"quiz_id": quiz.ID,
	})

	subject, err := u.subjectRepo.FindOneBy(map[string]interface{}{
		"id": quiz.SubjectID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to find subject: %v", err)
	}

	quizSubmissions, err := u.quizSubmissionRepo.FindBy(map[string]interface{}{
		"quiz_id": quiz.ID,
		"user_id": user.ID,
	}, 0, 0)

	if err != nil {
		return nil, fmt.Errorf("failed to find quizSubmissions: %v", err)
	}

	// Check if the quiz has been completed or not.
	var status response.QuizStatus
	quizSubmissionCount := len(quizSubmissions)
	quizSubmissionUUID := ""
	max_score := 0

	if quizSubmissionCount > 0 {
		status = response.SudahDikerjakan
		// Check if it has been completed once or more
		if quizSubmissionCount == 1 {
			max_score = quizSubmissions[quizSubmissionCount-1].Score
			quizSubmissionUUID = quizSubmissions[quizSubmissionCount-1].UUID
		} else {
			// Find the quizSubmissionUUID with the highest value
			for _, quizSubmission := range quizSubmissions {
				if quizSubmission.Score > max_score {
					max_score = quizSubmission.Score
					quizSubmissionUUID = quizSubmission.UUID
				}
			}
		}

	} else {
		max_score = -1
		status = response.BelumDikerjakan
	}

	result := &response.QuizDetailResponse{
		ID:                 quiz.ID,
		UUID:               quiz.UUID,
		Subject:            subject.Name,
		Title:              quiz.Title,
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
