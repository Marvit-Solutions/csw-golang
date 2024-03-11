package mentor

import (
	"net/http"

	dto "github.com/Marvit-Solutions/csw-golang/internal/domain/response"

	"github.com/gin-gonic/gin"
)

func (mc *mentorHandler) ListThreeMentors(c *gin.Context) {
	data, err := mc.mentorUsecase.GetListTopThreeMentors()
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
		Data:    data,
	})

}

func (mc *mentorHandler) GetAllMentors(c *gin.Context) {
	data, err := mc.mentorUsecase.GetAllMentors()
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
		Data:    data,
	})

}
