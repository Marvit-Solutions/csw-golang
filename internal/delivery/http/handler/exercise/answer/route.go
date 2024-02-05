package answer

import "github.com/gin-gonic/gin"

func (exerciseAnswerHandler *ExerciseAnswerHandler) RegisterRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("/exercise-answer")
	authGroup.POST("/add", exerciseAnswerHandler.AddAnswer)
}
