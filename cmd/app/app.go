package app

import (
	db "csw-golang/pkg/database/psql"

	"csw-golang/internal/delivery"
	"csw-golang/internal/delivery/http/handler"
	authHandler "csw-golang/internal/delivery/http/handler/auth"
	mentorHandler "csw-golang/internal/delivery/http/handler/mentor"
	paketHandler "csw-golang/internal/delivery/http/handler/paket"
	authRepo "csw-golang/internal/domain/repository/auth"
	mentorRepo "csw-golang/internal/domain/repository/mentor"
	paketRepo "csw-golang/internal/domain/repository/paket"
	authUsecase "csw-golang/internal/domain/usecase/auth"
	mentorUsecase "csw-golang/internal/domain/usecase/mentor"
	paketUsecase "csw-golang/internal/domain/usecase/paket"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	db.Init()
	authRepo := authRepo.New(db.DB)
	authUsecase := authUsecase.New(authRepo)
	authHandler := authHandler.New(authUsecase)

	paketRepo := paketRepo.New(db.DB)
	paketUsecase := paketUsecase.New(paketRepo)
	paketHandler := paketHandler.New(paketUsecase)

	mentorRepo := mentorRepo.New(db.DB)
	mentorUsecase := mentorUsecase.New(mentorRepo)
	mentorHandler := mentorHandler.New(mentorUsecase)

	handler := handler.Handler{
		AuthHandler:   authHandler,
		PaketHandler:  paketHandler,
		MentorHandler: mentorHandler,
	}

	router := delivery.StartRoute(handler)
	return router
}
