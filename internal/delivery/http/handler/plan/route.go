package plan

import (
	"github.com/gin-gonic/gin"
)

func (planHandler *PlanHandler) RegisterRoutes(r *gin.RouterGroup) {
	planGroup := r.Group("/plan")
	planGroup.GET("/all", planHandler.ListPlan)
	planGroup.GET("/top", planHandler.GetTop3Plan)

}
