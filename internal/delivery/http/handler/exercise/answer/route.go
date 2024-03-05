package answer

import "github.com/gin-gonic/gin"

func (exerciseAnswerHandler *ExerciseAnswerHandler) RegisterRoutes(r *gin.RouterGroup) {
	exerciseAnswerGroup := r.Group("/exercise-answer")
	exerciseAnswerGroup.POST("/add", exerciseAnswerHandler.AddAnswer)
}
