package materialAdmin

import (
	materialAdmin "github.com/Marvit-Solutions/csw-golang/admin_module/app/usecase/material"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	u materialAdmin.Usecase
}

type Handler interface {
	MaterialAdminAll(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	MaterialUpdateDetail(ctx *gin.Context)
}

func NewHandler(
	db *gorm.DB,
) Handler {
	return &handler{
		u: materialAdmin.NewUsecase(db),
	}
}
