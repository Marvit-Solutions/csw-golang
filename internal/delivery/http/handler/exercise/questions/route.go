package questions

import "github.com/gin-gonic/gin"

func (exerciseQuestionsHandler *ExerciseQuestionsHandler) RegisterRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("/exercise-questions")
	authGroup.POST("/add", exerciseQuestionsHandler.AddExerciseQuestion)
	authGroup.POST("/addall", exerciseQuestionsHandler.AddBatchExerciseQuestion)
	authGroup.GET("/id", exerciseQuestionsHandler.GetExerciseQuestions)
	authGroup.GET("/all", exerciseQuestionsHandler.GetAllExerciseQuestions)

}
