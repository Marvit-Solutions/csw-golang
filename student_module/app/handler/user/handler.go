package user

import (
	"github.com/Marvit-Solutions/csw-golang/student_module/app/usecase/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	u user.Usecase
}

type Handler interface {
	Find(c *gin.Context)
}

func NewHandler(
	db *gorm.DB,
) Handler {
	return &handler{
		u: user.NewUsecase(db),
	}
}
