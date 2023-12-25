package app

import (
	"csw-golang/pkg/database"

	"csw-golang/internal/delivery"
	"csw-golang/internal/delivery/http/handler"
	authHandler "csw-golang/internal/delivery/http/handler/auth"
	paketHandler "csw-golang/internal/delivery/http/handler/paket"
	authRepo "csw-golang/internal/domain/repository/auth"
	paketRepo "csw-golang/internal/domain/repository/paket"
	authUsecase "csw-golang/internal/domain/usecase/auth"
	paketUsecase "csw-golang/internal/domain/usecase/paket"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	database.Init()

	authRepo := authRepo.New(database.DB)
	authUsecase := authUsecase.New(authRepo)
	authHandler := authHandler.New(authUsecase)

	paketRepo := paketRepo.New(database.DB)
	paketUsecase := paketUsecase.New(paketRepo)
	paketHandler := paketHandler.New(paketUsecase)

	handler := handler.Handler{
		AuthHandler:  authHandler,
		PaketHandler: paketHandler,
	}

	router := delivery.StartRoute(handler)
	return router
}
