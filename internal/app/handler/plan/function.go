package plan

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/library/helper"

	"github.com/gin-gonic/gin"
)

func (ph *planHandler) ListPlan(c *gin.Context) {
	data, err := ph.planUsecase.ListPlan()
	if err != nil {
		helper.NewErrorResponse(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	if len(data) == 0 {
		helper.NewErrorResponse(c, http.StatusNotFound, http.StatusText(http.StatusNotFound), err)
		return
	}

	helper.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data)
}

func (ph *planHandler) GetTop3Plan(c *gin.Context) {
	data, err := ph.planUsecase.GetTop3Plan()
	if err != nil {
		helper.NewErrorResponse(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	if len(data) == 0 {
		helper.NewErrorResponse(c, http.StatusNotFound, http.StatusText(http.StatusNotFound), err)
		return
	}

	helper.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data)
}
