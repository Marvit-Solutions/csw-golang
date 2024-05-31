package dashboard

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/usecase/dashboard"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	u dashboard.Usecase
}

type Handler interface {
	FindMaterial(c *gin.Context)
	QuizAll(c *gin.Context)
}

func NewHandler(
	db *gorm.DB,
) Handler {
	return &handler{
		u: dashboard.NewUsecase(db),
	}
}
