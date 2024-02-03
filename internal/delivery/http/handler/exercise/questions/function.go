package questions

import (
	"csw-golang/internal/domain/entity/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (mc *ExerciseQuestionsHandler) AddExerciseQuestion(c *gin.Context) {

	request := dto.QuestionExercisesRequest{}

	c.Bind(&request)

	question, err := mc.exerciseQuestionsUsecase.AddExerciseQuestion(request)
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

func (mc *ExerciseQuestionsHandler) GetExerciseQuestions(c *gin.Context) {
	param := c.Query("question-id")
	response, err := mc.exerciseQuestionsUsecase.GetExerciseQuestions(param)
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
		Data:    response,
	})

}
