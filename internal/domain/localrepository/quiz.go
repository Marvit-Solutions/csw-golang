package localrepository

import (
	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/response"
)

type Quiz interface {
	FindUserAnswerReview(quizSubmissionID int) ([]*response.UserAnswer, error)
	CountQuizzesGroupedBySubModule(moduleID int, testTypeID int) ([]*response.QuizzesGroupedBySubModule, error)
	GetQuizAll(testTypeID int, subModulID int, limit int, offset int) ([]*response.QuizItemAll, error)
	CountQuizALL(testTypeID int, subModulID int) (int, error)
}
