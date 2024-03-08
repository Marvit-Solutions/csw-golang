package module

import (
	"csw-golang/internal/domain/entity/request"
	"csw-golang/internal/domain/helper/response"
	"csw-golang/internal/domain/helper/validator"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (mc *ModuleHandler) GetListModules(c *gin.Context) {
	data, err := mc.moduleUsecase.GetListModules()

	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	if data == nil {
		response.NewErrorResponse(c, http.StatusNotFound, http.StatusText(http.StatusNotFound), nil)
		return
	}

	response.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data, "Success Get List Of Modules")

}

func (mc *ModuleHandler) GetSubjectsBySubmoduleID(c *gin.Context) {
	submoduleID := c.Param("submoduleID")

	data, err := mc.moduleUsecase.GetSubjectsBySubmoduleID(submoduleID)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	if data == nil {
		response.NewErrorResponse(c, http.StatusNotFound, http.StatusText(http.StatusNotFound), nil)
		return
	}

	response.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data, "Success Get Subjects By Submodule ID")
}

func (mc *ModuleHandler) GetQuestionsByTestTypeID(c *gin.Context) {
	testTypeID := c.Param("testTypeID")

	data, err := mc.moduleUsecase.GetQuestionsByTestTypeID(testTypeID)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	if data.QuestionTotal == 0 {
		response.NewErrorResponse(c, http.StatusNotFound, http.StatusText(http.StatusNotFound), nil)
		return
	}

	response.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data, "Success Get Questions By Test Type ID")
}

func (mc *ModuleHandler) GetTestReview(c *gin.Context) {
	submissionID := c.Param("submissionID")

	data, err := mc.moduleUsecase.GetTestReview(submissionID)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	if data.QuestionTotal == 0 {
		response.NewErrorResponse(c, http.StatusNotFound, http.StatusText(http.StatusNotFound), nil)
		return
	}

	response.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data, "Success Get Test Review")
}

func (mc *ModuleHandler) PostSubmittedTest(c *gin.Context) {

	testTypeID := c.Param("testTypeID")
	userID := "7aa65bf1-9273-46ad-ba2c-bf94ddcfcc6e"

	// userID, err := jwt.GetAuthenticatedUser(c.Request)
	// if err != nil {
	// 	response.NewErrorResponse(c, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), err)
	// 	return
	// }

	// userIDString := strconv.Itoa(userID)
	// fmt.Println("USER ID: ", userID)
	// fmt.Println("USER ID CONVERTED: ", userIDString)

	request := request.UserSubmittedQuizRequest{}

	if err := validator.BindingValidation(c, &request); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	err := mc.moduleUsecase.PostSubmittedTest(testTypeID, userID, request)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	response.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), nil, "Success Post Submitted Test")
}

func (mc *ModuleHandler) GetTop3EverySubject(c *gin.Context) {
	userID := "7aa65bf1-9273-46ad-ba2c-bf94ddcfcc6e"
	subjectTypeID := c.Param("subjectTypeID")

	fmt.Println("USER ID: ", userID)

	data, err := mc.moduleUsecase.GetTop3EverySubject(userID, subjectTypeID)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	if data == nil {
		response.NewErrorResponse(c, http.StatusNotFound, http.StatusText(http.StatusNotFound), nil)
		return
	}

	response.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data, "Success Get Top 3 Every Subject")

}
