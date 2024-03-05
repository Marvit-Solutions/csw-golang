package submission

import (
	"csw-golang/internal/domain/entity/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ms *SubmissionHandler) AddSubmission(c *gin.Context) {

	testTypeID := c.Query("testTypeID")
	userID := "7aa65bf1-9273-46ad-ba2c-bf94ddcfcc6e"
	request := dto.UserSubmittedAnswerExercisesRequest{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}
	question, err := ms.submissionUsecase.AddSubmissionUsecase(testTypeID, userID, request)
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
		Data:    question,
	})

}

func (ms *SubmissionHandler) GetSubmission(c *gin.Context) {

	value := c.Query("exercise-submission-id")
	question, err := ms.submissionUsecase.GetSubmissionUsecase(value)
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
		Data:    question,
	})

}

func (ms *SubmissionHandler) GetAllSubmission(c *gin.Context) {

	question, err := ms.submissionUsecase.GetAllSubmissionsUsecase()
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
		Data:    question,
	})

}
