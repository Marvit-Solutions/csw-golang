package plan

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/usecase/plan"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type planHandler struct {
	planUsecase plan.PlanUsecase
}

type PlanHandler interface {
	ListPlan(c *gin.Context)
	GetTop3Plan(c *gin.Context)
}

func NewPlanHandler(
	db *gorm.DB,
) PlanHandler {
	return &planHandler{
		planUsecase: plan.NewPlanUsecase(db),
	}
}
