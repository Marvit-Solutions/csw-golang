package quiz

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/usecase/quiz"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	u quiz.Usecase
}

type Handler interface {
	QuizContent(c *gin.Context)
	QuizSubmission(c *gin.Context)
	QuizDetail(c *gin.Context)
	QuizReview(c *gin.Context)
}

func NewHandler(
	db *gorm.DB,
) Handler {
	return &handler{
		u: quiz.NewUsecase(db),
	}
}
