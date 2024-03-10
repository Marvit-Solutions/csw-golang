package request

import (
	"csw-golang/library/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouteInit struct {
	Engine    *gin.Engine
	SQLMaster *gorm.DB
	Env       config.Config
}
