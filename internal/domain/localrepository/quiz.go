package localrepository

import "github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"

type Quiz interface {
	FindUserAnswerReview(quizSubmissionID int) ([]*response.UserAnswer, error)
	CountQuizzesGroupedBySubModule(moduleID int, testTypeID int) ([]*response.QuizzesGroupedBySubModule, error)
}
