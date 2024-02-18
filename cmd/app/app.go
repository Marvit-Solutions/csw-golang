package app

import (
	db "csw-golang/pkg/database/psql"

	"csw-golang/internal/delivery"
	"csw-golang/internal/delivery/http/handler"

	authHandler "csw-golang/internal/delivery/http/handler/auth"
	authRepo "csw-golang/internal/domain/repository/auth"
	authUsecase "csw-golang/internal/domain/usecase/auth"

	pretestHandler "csw-golang/internal/delivery/http/handler/pretest"
	pretestRepo "csw-golang/internal/domain/repository/pretest"
	pretestUsecase "csw-golang/internal/domain/usecase/pretest"

	mentorHandler "csw-golang/internal/delivery/http/handler/mentor"
	mentorRepo "csw-golang/internal/domain/repository/mentor"
	mentorUsecase "csw-golang/internal/domain/usecase/mentor"

	planHandler "csw-golang/internal/delivery/http/handler/plan"
	planRepo "csw-golang/internal/domain/repository/plan"
	planUsecase "csw-golang/internal/domain/usecase/plan"

	testimonialsHandler "csw-golang/internal/delivery/http/handler/testimonial"
	testimonialsRepo "csw-golang/internal/domain/repository/testimonial"
	testimonialUsecase "csw-golang/internal/domain/usecase/testimonial"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	db.Init()
	authRepo := authRepo.New(db.DB)
	authUsecase := authUsecase.New(authRepo)
	authHandler := authHandler.New(authUsecase)

	planRepo := planRepo.New(db.DB)
	planUsecase := planUsecase.New(planRepo)
	planHandler := planHandler.New(planUsecase)

	mentorRepo := mentorRepo.New(db.DB)
	mentorUsecase := mentorUsecase.New(mentorRepo)
	mentorHandler := mentorHandler.New(mentorUsecase)

	testimonialRepo := testimonialsRepo.New(db.DB)
	testimonialUsecase := testimonialUsecase.New(testimonialRepo)
	testimonialHandler := testimonialsHandler.New(testimonialUsecase)

	pretestRepo := pretestRepo.New(db.DB)
	pretestUsecase := pretestUsecase.New(pretestRepo)
	pretestHandler := pretestHandler.New(pretestUsecase)

	handler := handler.Handler{
		AuthHandler:        authHandler,
		PlanHandler:        planHandler,
		MentorHandler:      mentorHandler,
		TestimonialHandler: testimonialHandler,
		PretestHandler:     pretestHandler,
	}

	router := delivery.StartRoute(handler)
	return router
}
