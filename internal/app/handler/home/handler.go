package home

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/usecase/home"
	"github.com/Marvit-Solutions/csw-golang/library/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	u home.Usecase
}

type Handler interface {
	TopMentor(c *gin.Context)
	AllMentor(c *gin.Context)
	// TopPlan(c *gin.Context)
	// AllPlan(c *gin.Context)
}

func NewHandler(
	db *gorm.DB,
	conf config.Config,
) Handler {
	return &handler{
		u: home.NewUsecase(db, conf),
	}
}
