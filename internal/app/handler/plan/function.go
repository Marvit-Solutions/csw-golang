package plan

import (
	"net/http"

	dto "github.com/Marvit-Solutions/csw-golang/internal/domain/response"

	"github.com/gin-gonic/gin"
)

func (ph *planHandler) ListPlan(c *gin.Context) {

	data, err := ph.planUsecase.ListPlan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Code:    http.StatusInternalServerError,
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	if len(data) == 0 {
		c.JSON(http.StatusNotFound, dto.Success[interface{}]{
			Message: "Plans not found!",
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Success get plans!",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    data,
	})
}

func (ph *planHandler) GetTop3Plan(c *gin.Context) {
	data, err := ph.planUsecase.GetTop3Plan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if len(data) == 0 {
		c.JSON(http.StatusNotFound, dto.Success[interface{}]{
			Message: "Plans not found!",
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Success get top 3 plans!",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    data,
	})
}
