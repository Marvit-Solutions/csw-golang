package app

import (
	db "csw-golang/pkg/database/psql"

	"csw-golang/internal/delivery"
	"csw-golang/internal/delivery/http/handler"
	authHandler "csw-golang/internal/delivery/http/handler/auth"
	authRepo "csw-golang/internal/domain/repository/auth"
	authUsecase "csw-golang/internal/domain/usecase/auth"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	db.Init()
	authRepo := authRepo.New(db.DB)
	authUsecase := authUsecase.New(authRepo)
	authHandler := authHandler.New(authUsecase)

	handler := handler.Handler{
		AuthHandler: authHandler,
	}

	router := delivery.StartRoute(handler)
	return router
}
