package home

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/gin-gonic/gin"
)

func (h *handler) MentorAll(ctx *gin.Context) {

	mentors, err := h.u.MentorAll()
	if err != nil && err == helper.ErrDataNotFound {
		helper.NewErrorResponse(ctx, http.StatusNotFound, http.StatusText(http.StatusNotFound), err.Error())
		return
	}

	if err != nil {
		helper.NewErrorResponse(ctx, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), err.Error())
		return
	}

	helper.NewSuccessResponseNonPaged(ctx, http.StatusOK, http.StatusText(http.StatusOK), mentors)
}
