package auth

import (
	"csw-golang/internal/domain/entity/request"
	"csw-golang/internal/domain/helper/response"
	"csw-golang/internal/domain/helper/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ah *AuthHandler) Register(c *gin.Context) {
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

func (ah *AuthHandler) Login(c *gin.Context) {
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
