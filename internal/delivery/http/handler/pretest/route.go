package pretest

import "github.com/gin-gonic/gin"

func (pretestHandler *PretestHandler) RegisterRoutes(r *gin.RouterGroup) {
	testimonialGroup := r.Group("/pretest")
	testimonialGroup.GET("/all", pretestHandler.GetAllPretests)
	testimonialGroup.GET("/:id/soal", pretestHandler.GetPretestById)
}
