package home

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/gin-gonic/gin"
)

func (h *handler) TopMentor(ctx *gin.Context) {

	mentors, err := h.u.TopMentor()
	if err != nil {
		helper.NewErrorResponse(ctx, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	helper.NewSuccessResponseNonPaged(ctx, http.StatusOK, http.StatusText(http.StatusOK), mentors)
}

func (h *handler) AllMentor(ctx *gin.Context) {

	mentors, err := h.u.AllMentor()
	if err != nil {
		helper.NewErrorResponse(ctx, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	helper.NewSuccessResponseNonPaged(ctx, http.StatusOK, http.StatusText(http.StatusOK), mentors)
}
