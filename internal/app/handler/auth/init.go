package auth

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/usecase/auth"
	"github.com/Marvit-Solutions/csw-golang/library/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	u auth.Usecase
}

type Handler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

func NewHandler(
	db *gorm.DB,
	conf config.Config,
) Handler {
	return &handler{
		u: auth.NewUsecase(db, conf),
	}
}
