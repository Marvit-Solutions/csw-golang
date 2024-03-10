package mentor

import (
	"csw-golang/internal/domain/entity/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (mc *MentorHandler) ListThreeMentors(c *gin.Context) {
	err, response := mc.mentorUsecase.GetListTopThreeMentors()
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

func (mc *MentorHandler) GetAllMentors(c *gin.Context) {
	err, response := mc.mentorUsecase.GetAllMentors()
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
