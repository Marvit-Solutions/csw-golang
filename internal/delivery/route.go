package delivery

import (
	"csw-golang/internal/delivery/http/handler"
	"csw-golang/pkg/logger"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func StartRoute(handler handler.Handler) *gin.Engine {
	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}

	r.Use(cors.New(config))
	r.Use(logger.LogMiddleware())

	v1 := r.Group("/api/v1")
	{
		handler.AuthHandler.RegisterRoutes(v1)
		handler.PlanHandler.RegisterRoutes(v1)
		handler.MentorHandler.RegisterRoutes(v1)
		handler.TestimonialHandler.RegisterRoutes(v1)
		handler.PretestHandler.RegisterRoutes(v1)
	}

	return r
}
