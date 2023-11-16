package delivery

import (
	"csw-golang/internal/delivery/http/handler"
	"csw-golang/pkg/logger"

	"github.com/gin-gonic/gin"
)

func StartRoute(handler handler.Handler) *gin.Engine {
	r := gin.New()
	r.Use(logger.LogMiddleware())

	handler.AuthHandler.RegisterRoutes(r)

	return r
}
