package request

import (
	"github.com/Marvit-Solutions/csw-golang/library/config"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouteInit struct {
	Engine    *gin.Engine
	SQLMaster *gorm.DB
	Mongo     *mongo.Database
	Env       config.Config
}
