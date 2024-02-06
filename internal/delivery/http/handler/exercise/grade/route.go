package grade

import "github.com/gin-gonic/gin"

func (exerciseGradeHandler *ExerciseGradeHandler) RegisterRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("/exercise-grade")
	authGroup.POST("/add", exerciseGradeHandler.AddGrade)
	authGroup.GET("/get", exerciseGradeHandler.GetGrade)
	authGroup.GET("/id", exerciseGradeHandler.GetGradeById)
}
