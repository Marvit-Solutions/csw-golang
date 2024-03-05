package grade

import "github.com/gin-gonic/gin"

func (exerciseGradeHandler *ExerciseGradeHandler) RegisterRoutes(r *gin.RouterGroup) {
	exerciseGradeGroup := r.Group("/exercise-grade")
	exerciseGradeGroup.POST("/add", exerciseGradeHandler.AddGrade)
	exerciseGradeGroup.GET("/get", exerciseGradeHandler.GetGrade)
	exerciseGradeGroup.GET("/id", exerciseGradeHandler.GetGradeById)
}
