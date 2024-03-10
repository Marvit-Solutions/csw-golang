package auth

import (
	"csw-golang/internal/domain/usecase/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type authHandler struct {
	authUsecase auth.AuthUsecase
}

type AuthHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

func NewAuthHandler(
	db *gorm.DB,
) AuthHandler {
	return &authHandler{
		authUsecase: auth.NewAuthUsecase(db),
	}
}
