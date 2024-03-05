package questions

import "github.com/gin-gonic/gin"

func (exerciseQuestionsHandler *ExerciseQuestionsHandler) RegisterRoutes(r *gin.RouterGroup) {
	exerciseQuestionsGroup := r.Group("/exercise-questions")
	exerciseQuestionsGroup.POST("/add", exerciseQuestionsHandler.AddExerciseQuestion)
	exerciseQuestionsGroup.POST("/addall", exerciseQuestionsHandler.AddBatchExerciseQuestion)
	exerciseQuestionsGroup.GET("/id", exerciseQuestionsHandler.GetExerciseQuestions)
	exerciseQuestionsGroup.GET("/all", exerciseQuestionsHandler.GetAllExerciseQuestions)

}
