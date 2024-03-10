package auth

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/internal/domain/request"
	"github.com/Marvit-Solutions/csw-golang/library/helper/response"
	"github.com/Marvit-Solutions/csw-golang/library/helper/validator"

	"github.com/gin-gonic/gin"
)

func (ah *authHandler) Register(c *gin.Context) {
	var req request.RegisterRequest

	if err := validator.BindingValidation(c, &req); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	data, err := ah.authUsecase.Register(req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	response.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data, "Register Sukses!")
}

func (ah *authHandler) Login(c *gin.Context) {
	var req request.LoginRequest

	if err := validator.BindingValidation(c, &req); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	data, err := ah.authUsecase.Login(req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	response.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data, "Login Sukses!")
}
