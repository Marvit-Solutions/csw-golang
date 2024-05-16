package exercise

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/localmodel/request"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/gin-gonic/gin"
)

func (h *handler) FindDetail(c *gin.Context) {
	var req request.ExerciseDetailRequest

	if err := helper.ValidateURLParams(c, &req); err != nil {
		helper.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	exercise, err := h.u.FindDetail(req)
	if err != nil {
		helper.NewErrorResponse(c, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), err.Error())
		return
	}

	helper.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), exercise)
}
