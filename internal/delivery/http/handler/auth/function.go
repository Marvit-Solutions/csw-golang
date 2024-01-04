package auth

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/helper/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ah *AuthHandler) Register(c *gin.Context) {
	var request dto.RegisterRequest

	if err := validator.Validation(c, &request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
			"Code":    http.StatusBadRequest,
			"Status":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	err := ah.authUsecase.Register(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
			Status:  http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Register sukses!",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
	})
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var request dto.LoginRequest

	if err := validator.Validation(c, &request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
			"Code":    http.StatusBadRequest,
			"Status":  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	response, err := ah.authUsecase.Login(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Login sukses!",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    response,
	})
}
