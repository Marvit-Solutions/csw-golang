package grade

import (
	"csw-golang/internal/domain/entity/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (exerciseGradeHandler *ExerciseGradeHandler) AddGrade(c *gin.Context) {

	request := dto.GradeExercisesRequest{}

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}
	question, err := exerciseGradeHandler.exerciseAnswerUsecase.AddGradeUsecase(request)
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

func (exerciseGradeHandler *ExerciseGradeHandler) GetGrade(c *gin.Context) {

	userId := c.Query("user-id")
	submissionId := c.Query("submission-id")

	question, err := exerciseGradeHandler.exerciseAnswerUsecase.GetGradeUsecase(submissionId, userId)
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

func (exerciseGradeHandler *ExerciseGradeHandler) GetGradeById(c *gin.Context) {

	gradeID := c.Query("grade-id")

	question, err := exerciseGradeHandler.exerciseAnswerUsecase.GetGradeByIdUsecase(gradeID)
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
