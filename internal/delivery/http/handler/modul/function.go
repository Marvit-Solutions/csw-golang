package module

import (
	"csw-golang/internal/domain/entity/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (mc *ModuleHandler) GetListModules(c *gin.Context) {
	response, err := mc.moduleUsecase.GetListModules()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	if response == nil {
		c.JSON(http.StatusNotFound, dto.Fail{
			Message: "Data not found",
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Success",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    response,
	})

}

func (mc *ModuleHandler) GetSubjectsBySubmoduleID(c *gin.Context) {
	submoduleID := c.Param("id")

	response, err := mc.moduleUsecase.GetSubjectsBySubmoduleID(submoduleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	if response == nil {
		c.JSON(http.StatusNotFound, dto.Fail{
			Message: "Data not found",
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Success",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    response,
	})
}

func (mc *ModuleHandler) GetQuestionsByTestTypeID(c *gin.Context) {
	testTypeID := c.Param("id")

	response, err := mc.moduleUsecase.GetQuestionsByTestTypeID(testTypeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	if response.QuestionTotal == 0 {
		c.JSON(http.StatusNotFound, dto.Fail{
			Message: "Data not found",
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Success",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    response,
	})
}
