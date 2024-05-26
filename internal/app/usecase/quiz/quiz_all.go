package quiz

import (
	"fmt"
	"math"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
)

func (u *usecase) QuizAll(req request.ParamQuizAll) (*response.QuizAllResponse, int, int, error) {
	limit := req.Limit
	offset := 0

	if req.Page > 0 {
		offset = (req.Page - 1) * limit
	}

	// get the pretest that has been done
	subModule, err := u.subModulRepo.FindOneBy(map[string]interface{}{
		"uuid": req.SubModuleUUID,
	})

	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to find subModule: %v", err)
	}

	module, err := u.modulRepo.FindOneBy(map[string]interface{}{
		"id": subModule.ModuleID,
	})

	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to find module: %v", err)
	}

	countQuizAll, err := u.quizLocalRepo.CountQuizALL(req.TestTypeId, subModule.ID)
	totalRows := countQuizAll
	totalPages := int(math.Ceil(float64(countQuizAll) / float64(req.Limit)))
	fmt.Println(countQuizAll)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to countQuizAll: %v", err)
	}

	quizAll, err := u.quizLocalRepo.GetQuizAll(req.TestTypeId, subModule.ID, limit, offset)
	if err != nil {
		return nil, 0, 0, fmt.Errorf("failed to find quizAll: %v", err)
	}

	result := &response.QuizAllResponse{
		UUID:          req.SubModuleUUID,
		SubModuleName: subModule.Name,
		ModuleName:    module.Name,
		Quizzes:       quizAll,
	}

	// quizSubmission,err := u.quizSubmissionRepo.FindBy(map[string]interface{}{
	// 	"user_id": 40,

	// },0,0)

	// quiz, err := u.quizRepo.FindOneBy(map[string]interface{}{
	// 	"uuid": req.QuizUUID,
	// })
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to find subject: %v", err)
	// }

	// quizQuestionTotal := u.quizQuestionRepo.Count(map[string]interface{}{
	// 	"quiz_id": quiz.ID,
	// })

	// // find subject name
	// // tmpQuizSubjectID := "ce75821b-8641-4705-896c-5175c0fa9ce0"
	// subject, err := u.subjectRepo.FindOneBy(map[string]interface{}{
	// 	"id": quiz.SubjectID,
	// })
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to find subject: %v", err)
	// }

	// quizSubmissions, err := u.quizSubmissionRepo.FindBy(map[string]interface{}{
	// 	"quiz_id": quiz.ID,
	// 	"user_id": 40,
	// }, 0, 0)

	// if err != nil {
	// 	return nil, fmt.Errorf("failed to find subject: %v", err)
	// }

	// // cek sudah dikerjakan atau belum
	// var status response.QuizStatus
	// // quizSubmissionCount := u.quizSubmissionRepo.Count(map[string]interface{}{
	// // 	"quiz_id": quiz.ID,
	// // 	"user_id": req.UserID,
	// // })
	// quizSubmissionCount := len(quizSubmissions)
	// quizSubmissionUUID := ""
	// max_score := 0

	// if quizSubmissionCount > 0 {
	// 	status = response.SudahDikerjakan
	// 	// cek jika sudah dikerjakan sekali atau lebih
	// 	if quizSubmissionCount == 1 {
	// 		max_score = quizSubmissions[quizSubmissionCount-1].Score
	// 		quizSubmissionUUID = quizSubmissions[quizSubmissionCount-1].UUID
	// 	} else {
	// 		// cari quizSubmissionUUID dengan nilai yang paling tinggi
	// 		for _, quizSubmission := range quizSubmissions {
	// 			if quizSubmission.Score > max_score {
	// 				max_score = quizSubmission.Score
	// 				quizSubmissionUUID = quizSubmission.UUID
	// 			}
	// 		}
	// 	}

	// } else {
	// 	status = response.BelumDikerjakan
	// }

	// fmt.Println("ini quizSubmissionUUID")
	// fmt.Println(quizSubmissionUUID)

	// result := &response.QuizAllResponse{
	// 	ID:                 quiz.ID,
	// 	UUID:               quiz.UUID,
	// 	Subject:            subject.Name,
	// 	Modul:              quiz.Title,
	// 	Description:        quiz.Description,
	// 	TotalQuestions:     quizQuestionTotal,
	// 	TotalTime:          quiz.Time,
	// 	Status:             status,
	// 	AttemptAllowed:     quiz.Attempt,
	// 	QuizSubmissionUUID: quizSubmissionUUID,
	// 	Score:              max_score,
	// 	MaxScore:           quiz.MaxScore,
	// 	Attempt:            quizSubmissionCount,
	// }

	return result, totalRows, totalPages, nil

}
