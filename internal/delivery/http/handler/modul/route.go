package module

import "github.com/gin-gonic/gin"

func (moduleHandler *ModuleHandler) RegisterRoutes(r *gin.RouterGroup) {
	moduleGroup := r.Group("/module")
	moduleGroup.GET("/all", moduleHandler.GetListModules)
	moduleGroup.GET("/subjects/:id", moduleHandler.GetSubjectsBySubmoduleID)

}
