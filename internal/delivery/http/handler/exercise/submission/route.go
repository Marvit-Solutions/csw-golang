package submission

import "github.com/gin-gonic/gin"

func (submissionHandler *SubmissionHandler) RegisterRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("/exercise-submission")
	authGroup.POST("/add", submissionHandler.AddSubmission)
	authGroup.GET("/id", submissionHandler.GetSubmission)
	authGroup.GET("/all", submissionHandler.GetAllSubmission)
}
