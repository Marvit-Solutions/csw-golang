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
	MentorTop(c *gin.Context)
	MentorAll(c *gin.Context)
	MentorDetail(c *gin.Context)
	PlanTop(c *gin.Context)
	PlanAll(c *gin.Context)
	Testimonial(c *gin.Context)
}

func NewHandler(
	db *gorm.DB,
	conf config.Config,
) Handler {
	return &handler{
		u: home.NewUsecase(db, conf),
	}
}
