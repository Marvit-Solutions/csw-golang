package quizAdmin

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/usecase/quizAdmin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	u quizAdmin.Usecase
}

type Handler interface {
	QuizAdminAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	QuizUpdateDetail(c *gin.Context)
}

func NewHandler(
	db *gorm.DB,
) Handler {
	return &handler{
		u: quizAdmin.NewUsecase(db),
	}
}
