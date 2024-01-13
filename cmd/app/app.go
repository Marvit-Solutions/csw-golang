package app

import (
	db "csw-golang/pkg/database/psql"

	"csw-golang/internal/delivery"
	"csw-golang/internal/delivery/http/handler"
	authHandler "csw-golang/internal/delivery/http/handler/auth"
	mentorHandler "csw-golang/internal/delivery/http/handler/mentor"
	paketHandler "csw-golang/internal/delivery/http/handler/paket"
	testimonialsHandler "csw-golang/internal/delivery/http/handler/testimonial"

	authRepo "csw-golang/internal/domain/repository/auth"
	mentorRepo "csw-golang/internal/domain/repository/mentor"
	paketRepo "csw-golang/internal/domain/repository/paket"
	testimonialsRepo "csw-golang/internal/domain/repository/testimonial"

	authUsecase "csw-golang/internal/domain/usecase/auth"
	mentorUsecase "csw-golang/internal/domain/usecase/mentor"
	paketUsecase "csw-golang/internal/domain/usecase/paket"
	testimonialUsecase "csw-golang/internal/domain/usecase/testimonial"

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

	testimonialRepo := testimonialsRepo.New(db.DB)
	testimonialUsecase := testimonialUsecase.New(testimonialRepo)
	testimonialHandler := testimonialsHandler.New(testimonialUsecase)

	handler := handler.Handler{
		AuthHandler:        authHandler,
		PaketHandler:       paketHandler,
		MentorHandler:      mentorHandler,
		TestimonialHandler: testimonialHandler,
	}

	router := delivery.StartRoute(handler)
	return router
}
