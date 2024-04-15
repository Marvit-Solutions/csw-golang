package module

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/usecase/module"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	u module.Usecase
}

type Handler interface {
	ModuleAll(c *gin.Context)
	ModuleDetail(c *gin.Context)
	MaterialAll(c *gin.Context)
	MaterialFind(c *gin.Context)
}

func NewHandler(
	db *gorm.DB,
) Handler {
	return &handler{
		u: module.NewUsecase(db),
	}
}
