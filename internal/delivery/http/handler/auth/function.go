package auth

import (
	"csw-golang/internal/domain/entity/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ah *AuthHandler) Register(c *gin.Context) {
	var request dto.RegisterRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail{
			Status:  http.StatusText(http.StatusBadRequest),
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err := ah.authUsecase.Register(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Status:  http.StatusText(http.StatusInternalServerError),
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Status:  http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Message: "Register sukses!",
	})
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail{
			Status:  http.StatusText(http.StatusBadRequest),
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	response, err := ah.authUsecase.Login(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Status:  http.StatusText(http.StatusInternalServerError),
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Status:  http.StatusText(http.StatusOK),
		Code:    http.StatusOK,
		Message: "Login sukses!",
		Data:    response,
	})
}
