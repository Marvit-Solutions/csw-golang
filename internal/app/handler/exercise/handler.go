package exercise

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/usecase/exercise"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	u exercise.Usecase
}

type Handler interface {
	FindAll(c *gin.Context)
	FindDetail(c *gin.Context)
}

func NewHandler(
	db *gorm.DB,
) Handler {
	return &handler{
		u: exercise.NewUsecase(db),
	}
}
