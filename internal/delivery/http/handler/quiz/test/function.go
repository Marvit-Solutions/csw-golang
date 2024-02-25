package tests

import (
	"csw-golang/internal/domain/entity/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (th *TestHandler) GetAllTests(c *gin.Context) {

	response, err := th.testUsecase.GetAllTests()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Code:    http.StatusInternalServerError,
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	if len(response) == 0 {
		c.JSON(http.StatusNotFound, dto.Success[interface{}]{
			Message: "Test not found!",
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[[]dto.QuizResponse]{
		Message: "Success get test!",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    response,
	})
}
