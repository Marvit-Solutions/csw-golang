package auth

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/model/request"
	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/gin-gonic/gin"
)

func (h *handler) Login(ctx *gin.Context) {
	var req request.LoginRequest

	if err := helper.ValidateRequestBody(ctx, &req); err != nil {
		helper.NewErrorResponse(ctx, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	user, err := h.u.Login(req)
	if err != nil {
		helper.NewErrorResponse(ctx, http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), err.Error())
		return
	}

	helper.NewSuccessResponseNonPaged(ctx, http.StatusOK, http.StatusText(http.StatusOK), user)
}
