package mentor

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/library/helper"

	"github.com/gin-gonic/gin"
)

func (mc *mentorHandler) ListThreeMentors(c *gin.Context) {
	data, err := mc.mentorUsecase.GetListTopThreeMentors()
	if err != nil {
		helper.NewErrorResponse(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	helper.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data)
}

func (mc *mentorHandler) GetAllMentors(c *gin.Context) {
	data, err := mc.mentorUsecase.GetAllMentors()
	if err != nil {
		helper.NewErrorResponse(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	helper.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data)
}
