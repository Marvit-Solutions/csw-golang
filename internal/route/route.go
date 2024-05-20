package route

import (
	"github.com/Marvit-Solutions/csw-golang/internal"
	"github.com/Marvit-Solutions/csw-golang/library/middleware/cors"
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

	// Use CORS middleware.
	route.Use(cors.CORSMiddleware())

	// Use Gin's built-in logging middleware.
	route.Use(gin.Logger())

	// Use Gin's recovery middleware.
	route.Use(gin.Recovery())

	// Define route cors options.
	route.OPTIONS("/*path", cors.CORSMiddleware())

	// Define routes for different endpoints.
	// Auth routes
	{
		authGroup := route.Group("/auth")
		authGroup.POST("/register", module.Auth.Register)
		authGroup.POST("/login", module.Auth.Login)
	}

	// Home routes
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

	// Routes for get location
	{
		locationGroup := route.Group("/location")
		locationGroup.GET("/province", module.Location.Province)
		locationGroup.GET("/regency/:province_id", module.Location.Regency)
		locationGroup.GET("/district/:regency_id", module.Location.District)
	}

	// Routes for dashboard
	{
		dashboardGroup := route.Group("/dashboard")
		// Routes for student dashboard
		{
			dashboardStudentGroup := dashboardGroup.Group("/student")
			// Routes for module and material
			{
				modulGroup := dashboardStudentGroup.Group("/module")
				modulGroup.GET("/all", module.Module.ModuleAll)
				modulGroup.GET(":sub_module_uuid", module.Module.ModuleDetail)

				materiGroup := modulGroup.Group("/material")
				materiGroup.GET(":subject_uuid", module.Module.MaterialAll)
				materiGroup.GET("", module.Module.MaterialFind)
			}
			quizGroup := dashboardStudentGroup.Group("/quiz")
			quizGroup.GET("/quiz_content/:quiz_uuid", module.Quiz.QuizContent)
			quizGroup.POST("/quiz_submission", module.Quiz.QuizSubmission)
			quizGroup.GET("/quiz_detail/:quiz_uuid/:test_type_id", module.Quiz.QuizDetail)
			quizGroup.GET("/quiz_review/:quiz_submission_uuid/:quiz_uuid/:test_type_id", module.Quiz.QuizReview)
			quizGroup.GET("/quiz_score_all/:sub_modul_uuid", module.Quiz.QuizScoreAll)
			quizGroup.GET("/quiz_sub_module_all/:module_id/:test_type_id", module.Quiz.QuizSubModuleAll)

			// Routes for Exercise
			{
				exerciseGroup := dashboardStudentGroup.Group("/exercise")
				exerciseGroup.GET("all", module.Exercise.FindAll)
				exerciseGroup.GET(":exercise_uuid", module.Exercise.FindDetail)
			}
		}
	}
}
