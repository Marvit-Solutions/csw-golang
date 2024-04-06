package auth

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/request"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/gin-gonic/gin"
)

func (h *handler) Register(ctx *gin.Context) {
	var req request.RegisterRequest

	if err := helper.BindingValidation(ctx, &req); err != nil {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	user, err := h.u.Register(req)
	if err != nil {
		helper.NewErrorResponse(ctx, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	helper.NewSuccessResponseNonPaged(ctx, http.StatusOK, http.StatusText(http.StatusOK), user)
}
