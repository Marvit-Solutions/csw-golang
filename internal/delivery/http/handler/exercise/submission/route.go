package submission

import "github.com/gin-gonic/gin"

func (submissionHandler *SubmissionHandler) RegisterRoutes(r *gin.RouterGroup) {
	exerciseSubmissionGroup := r.Group("/exercise-submission")
	exerciseSubmissionGroup.POST("/add", submissionHandler.AddSubmission)
	exerciseSubmissionGroup.GET("/id", submissionHandler.GetSubmission)
	exerciseSubmissionGroup.GET("/all", submissionHandler.GetAllSubmission)
}
