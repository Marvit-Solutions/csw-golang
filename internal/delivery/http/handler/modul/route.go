package module

import "github.com/gin-gonic/gin"

func (moduleHandler *ModuleHandler) RegisterRoutes(r *gin.RouterGroup) {
	moduleGroup := r.Group("/module")
	moduleGroup.GET("/all", moduleHandler.GetListModules)
	moduleGroup.GET("/subjects/:submoduleID", moduleHandler.GetSubjectsBySubmoduleID)
	moduleGroup.GET("/exercise/:testTypeID/question", moduleHandler.GetQuestionsByTestTypeID)
	moduleGroup.GET("/exercise/review/:submissionID", moduleHandler.GetTestReview)
	moduleGroup.POST("/exercise/:testTypeID/question", moduleHandler.PostSubmittedTest)
	moduleGroup.GET("/history/top/:subjectTypeID", moduleHandler.GetTop3EverySubject)

}
