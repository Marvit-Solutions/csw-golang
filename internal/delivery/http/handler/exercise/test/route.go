package test

import "github.com/gin-gonic/gin"

func (exerciseTestHandler *ExerciseTestHandler) RegisterRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("/exercise-test")
	authGroup.POST("/add", exerciseTestHandler.AddExerciseTest)
	authGroup.GET("/id", exerciseTestHandler.GetExerciseTest)
	authGroup.GET("/all", exerciseTestHandler.GetAllExerciseTest)
}
