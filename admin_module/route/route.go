package route

import (
	"github.com/Marvit-Solutions/csw-golang/admin_module"
	"github.com/Marvit-Solutions/csw-golang/library/middleware/cors"
	"github.com/Marvit-Solutions/csw-golang/library/struct/request"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/module/apmgin"
)

// NewRouteInit initializes routes for various endpoints.
func NewRouteInit(req request.RouteInit) {
	// Create a new instance of dependencies using InitialInjection.
	module := admin_module.NewInitialInjection(req.SQLMaster, req.Env)

	// Create a new Gin route group.
	route := req.Engine.Group("admin-api/v1")

	// Add Elastic APM middleware to the route.
	route.Use(apmgin.Middleware(req.Engine))

	// Use CORS middleware.
	route.Use(cors.CORSMiddleware())

	// Use Gin's built-in logging middleware.
	route.Use(gin.Logger())

	// Use Gin's recovery middleware.
	route.Use(gin.Recovery())

	// Define route cors options.
	route.OPTIONS("/*path", cors.CORSMiddleware())

	// Define routes for different endpoints.
	Use(module)

	{
		quizGroup := route.Group("/quizzes")
		quizGroup.GET("", module.QuizAdmin.QuizAdminAll)
		quizGroup.GET("/update-details/:uuid", module.QuizAdmin.QuizUpdateDetail)
		quizGroup.POST("", module.QuizAdmin.Create)
		quizGroup.PUT("", module.QuizAdmin.Update)
		quizGroup.DELETE(":uuid", module.QuizAdmin.Delete)
	}

	{
		subjectGroup := route.Group("/subjects")
		subjectGroup.GET("", module.SubjectAdmin.SubjectAdminAll)
		subjectGroup.GET("/all", module.SubjectAdmin.Read)
		subjectGroup.GET("/update-details/:uuid", module.SubjectAdmin.SubjectUpdateDetail)
		subjectGroup.POST("", module.SubjectAdmin.Create)
		subjectGroup.PUT("", module.SubjectAdmin.Update)
		subjectGroup.DELETE(":uuid", module.SubjectAdmin.Delete)
	}

	{
		materialGroup := route.Group("/materials")
		materialGroup.GET("", module.MaterialAdmin.MaterialAdminAll)
		materialGroup.GET("/update-details/:uuid", module.MaterialAdmin.MaterialUpdateDetail)
		materialGroup.POST("", module.MaterialAdmin.Create)
		materialGroup.PUT("", module.MaterialAdmin.Update)
		materialGroup.DELETE(":uuid", module.MaterialAdmin.Delete)
	}
}

// Delete this function after "module" used
func Use(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}
