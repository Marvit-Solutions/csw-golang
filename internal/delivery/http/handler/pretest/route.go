package pretest

import "github.com/gin-gonic/gin"

func (pretestHandler *PretestHandler) RegisterRoutes(r *gin.RouterGroup) {
	pretestGroup := r.Group("/pretest")
	pretestGroup.GET("/all", pretestHandler.GetAllPretests)
	// pretestGroup.GET("/:id", pretestHandler.GetPretestById)
	// pretestGroup.GET("/:id/review", pretestHandler.GetPretestReview)
	// pretestGroup.POST("/:id/submit", pretestHandler.SubmitPretest)
}
