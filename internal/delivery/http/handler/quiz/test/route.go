package tests

import "github.com/gin-gonic/gin"

func (testHandler *TestHandler) RegisterRoutes(r *gin.RouterGroup) {
	quizGroup := r.Group("/quiz")
	quizGroup.GET("/all", testHandler.GetAllTests)
}
