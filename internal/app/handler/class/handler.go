package class

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/usecase/class"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	u class.Usecase
}

type Handler interface {
	FindAll(c *gin.Context)
}

func NewHandler(
	db *gorm.DB,
) Handler {
	return &handler{
		u: class.NewUsecase(db),
	}
}
