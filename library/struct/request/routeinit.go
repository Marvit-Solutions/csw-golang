package request

import (
	"github.com/Marvit-Solutions/csw-golang/library/config"
	"github.com/gomodule/redigo/redis"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouteInit struct {
	Engine    *gin.Engine
	SQLMaster *gorm.DB
	Redis     *redis.Pool
	Mongo     *mongo.Database
	Env       config.Config
}
