package answer

import (
	"csw-golang/internal/domain/entity/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (exerciseAnswerHandler *ExerciseAnswerHandler) AddAnswer(c *gin.Context) {

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
	question, err := exerciseAnswerHandler.exerciseAnswerUsecase.AddAnswerUsecase(request)
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
