package app

import (
	"csw-golang/cmd"
	"csw-golang/library/struct/request"

	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

func NewRouteInit(req request.RouteInit) {

	module := cmd.NewInitialInjection(req.SQLMaster, req.Env)
	route := req.Engine.Group("api/v1")
	route.Use(apmgin.Middleware(req.Engine))
	route.Use(gin.Logger())
	route.Use(gin.Recovery())

	{
		authGroup := route.Group("/account")
		authGroup.POST("/register", module.Auth.Register)
		authGroup.POST("/login", module.Auth.Login)
	}

	{
		mentorGroup := route.Group("/mentor")
		mentorGroup.GET("/top", module.Mentor.ListThreeMentors)
		mentorGroup.GET("/all", module.Mentor.GetAllMentors)
	}

	{
		planGroup := route.Group("/plan")
		planGroup.GET("/all", module.Plan.ListPlan)
		planGroup.GET("/top", module.Plan.GetTop3Plan)
	}
}
