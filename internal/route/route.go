package route

import (
	"github.com/Marvit-Solutions/csw-golang/internal"
	"github.com/Marvit-Solutions/csw-golang/library/struct/request"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

// NewRouteInit initializes routes for various endpoints.
func NewRouteInit(req request.RouteInit) {
	// Create a new instance of dependencies using InitialInjection.
	module := internal.NewInitialInjection(req.SQLMaster, req.Env)

	// Create a new Gin route group.
	route := req.Engine.Group("api/v1")

	// Add Elastic APM middleware to the route.
	route.Use(apmgin.Middleware(req.Engine))

	// Use Gin's built-in logging middleware.
	route.Use(gin.Logger())

	// Use Gin's recovery middleware.
	route.Use(gin.Recovery())

	// Define routes for different endpoints.
	{
		authGroup := route.Group("/auth")
		authGroup.POST("/register", module.Auth.Register)
		authGroup.POST("/login", module.Auth.Login)
	}

	{
		homeGroup := route.Group("/home")

		mentorGroup := homeGroup.Group("/mentor")
		mentorGroup.GET("top", module.Home.MentorTop)
		mentorGroup.GET("all", module.Home.MentorAll)
		mentorGroup.GET(":uuid", module.Home.MentorDetail)

		planGroup := homeGroup.Group("/plan")
		planGroup.GET("all", module.Home.PlanAll)
		planGroup.GET("top", module.Home.PlanTop)

		testimonialGroup := homeGroup.Group("/testimonial")
		testimonialGroup.GET("all", module.Home.Testimonial)
	}

	{
		locationGroup := route.Group("/location")
		locationGroup.GET("/province", module.Location.Province)
		locationGroup.GET("/regency/:province_id", module.Location.Regency)
		locationGroup.GET("/district/:regency_id", module.Location.District)
	}

	{
		modulGroup := route.Group("/modul")
		modulGroup.GET("/all", module.Modul.ModuleAll)
		modulGroup.GET(":sub_module_uuid", module.Modul.ModuleDetail)

		materiGroup := modulGroup.Group("/materi")
		materiGroup.GET(":subject_uuid", module.Modul.MaterialAll)
		materiGroup.GET("", module.Modul.MaterialFind)
	}
}
