package paket

import (
	"github.com/gin-gonic/gin"
)

func (paketHandler *PlanHandler) RegisterRoutes(r *gin.RouterGroup) {
	paketGroup := r.Group("/paket")
	paketGroup.GET("/all", paketHandler.ListPlan)
	paketGroup.GET("/top", paketHandler.GetTop3Plan)
	// paketGroup.POST("/", paketHandler.CreatePaket)
	// paketGroup.PUT("/:id_paket", paketHandler.UpdatePaket)
	// paketGroup.DELETE("/:id_paket", paketHandler.DeletePaket)

	// SubPlanGroup := r.Group("/SubPlan")
	// SubPlanGroup.GET("/:id_paket", paketHandler.ListSubPlan)
	// SubPlanGroup.POST("/", paketHandler.CreateSubPlan)
	// SubPlanGroup.PUT("/:id_SubPlan", paketHandler.UpdateSubPlan)
	// SubPlanGroup.DELETE("/:id_SubPlan", paketHandler.DeleteSubPlan)
	// SubPlanGroup.GET("/top", paketHandler.GetTopSubPlan)

}
