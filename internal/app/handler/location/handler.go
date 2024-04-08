package location

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/usecase/location"
	"github.com/Marvit-Solutions/csw-golang/library/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	u    location.Usecase
	conf config.Config
}

type Handler interface {
	Province(c *gin.Context)
	// Regency(c *gin.Context)
	// District(c *gin.Context)
}

func NewHandler(
	db *gorm.DB,
	conf config.Config,
) Handler {
	return &handler{
		u:    location.NewUsecase(db),
		conf: conf,
	}
}
