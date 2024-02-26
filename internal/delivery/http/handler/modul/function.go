package module

import (
	"csw-golang/internal/domain/entity/dto"
	"fmt"
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
	submoduleID := c.Param("submoduleID")

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
	testTypeID := c.Param("testTypeID")

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

func (mc *ModuleHandler) GetTestReview(c *gin.Context) {
	submissionID := c.Param("submissionID")
	// status := c.Param("status")

	response, err := mc.moduleUsecase.GetTestReview(submissionID)
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

func (mc *ModuleHandler) PostSubmittedTest(c *gin.Context) {

	testTypeID := c.Param("testTypeID")
	userID := "7aa65bf1-9273-46ad-ba2c-bf94ddcfcc6e"
	request := dto.UserSubmittedQuizRequest{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	err = mc.moduleUsecase.PostSubmittedTest(testTypeID, userID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Success",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    nil,
	})
}

func (mc *ModuleHandler) GetTop3EverySubject(c *gin.Context) {
	userID := "7aa65bf1-9273-46ad-ba2c-bf94ddcfcc6e"
	subjectTypeID := c.Param("subjectTypeID")

	fmt.Println("USER ID: ", userID)

	response, err := mc.moduleUsecase.GetTop3EverySubject(userID, subjectTypeID)
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
