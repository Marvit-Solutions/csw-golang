package auth

import (
	"github.com/gin-gonic/gin"
)

func (authHandler *AuthHandler) RegisterRoutes(r *gin.Engine) {
	authGroup := r.Group("/account")
	authGroup.POST("/register", authHandler.Register)
	authGroup.POST("/login", authHandler.Login)
}
