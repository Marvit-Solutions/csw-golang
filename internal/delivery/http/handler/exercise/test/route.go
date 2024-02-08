package test

import "github.com/gin-gonic/gin"

func (exerciseTestHandler *ExerciseTestHandler) RegisterRoutes(r *gin.RouterGroup) {
	exerciseTestGroup := r.Group("/exercise-test")
	exerciseTestGroup.POST("/add", exerciseTestHandler.AddExerciseTest)
	exerciseTestGroup.GET("/id", exerciseTestHandler.GetExerciseTest)
	exerciseTestGroup.GET("/all", exerciseTestHandler.GetAllExerciseTest)
}
