package auth

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/request"
	"github.com/Marvit-Solutions/csw-golang/library/helper"

	"github.com/gin-gonic/gin"
)

func (ah *authHandler) Register(c *gin.Context) {
	var req request.RegisterRequest

	if err := helper.BindingValidation(c, &req); err != nil {
		helper.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	data, err := ah.authUsecase.Register(req)
	if err != nil {
		helper.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	helper.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data)
}

func (ah *authHandler) Login(c *gin.Context) {
	var req request.LoginRequest

	if err := helper.BindingValidation(c, &req); err != nil {
		helper.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	data, err := ah.authUsecase.Login(req)
	if err != nil {
		helper.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	helper.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data)
}
