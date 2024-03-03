package auth

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/helper/response"
	"csw-golang/internal/domain/helper/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ah *AuthHandler) Register(c *gin.Context) {
	var request dto.RegisterRequest

	if err := validator.Validation(c, &request); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	err := ah.authUsecase.Register(request)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	response.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), nil, "Register Sukses!")
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var request dto.LoginRequest

	if err := validator.Validation(c, &request); err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	data, err := ah.authUsecase.Login(request)
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err)
		return
	}

	response.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data, "Login Sukses!")
}
