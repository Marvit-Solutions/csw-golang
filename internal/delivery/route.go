package delivery

import (
	"csw-golang/internal/delivery/http/handler"
	"csw-golang/pkg/logger"

	"github.com/gin-gonic/gin"
)

func StartRoute(handler handler.Handler) *gin.Engine {
	r := gin.New()
	r.Use(logger.LogMiddleware())

	v1 := r.Group("/api/v1")
	{
		handler.AuthHandler.RegisterRoutes(v1)
		handler.PaketHandler.RegisterRoutes(v1)
	}

	return r
}
