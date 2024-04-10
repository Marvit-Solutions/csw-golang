package modul

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/usecase/modul"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	u modul.Usecase
}

type Handler interface {
	ModulAll(c *gin.Context)
}

func NewHandler(
	db *gorm.DB,
) Handler {
	return &handler{
		u: modul.NewUsecase(db),
	}
}
