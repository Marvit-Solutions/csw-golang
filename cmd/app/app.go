package app

import (
	db "csw-golang/pkg/database/psql"
	"fmt"

	"csw-golang/internal/delivery"
	"csw-golang/internal/delivery/http/handler"
	"csw-golang/internal/delivery/http/middleware/jwt"

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

	quizTestHandler "csw-golang/internal/delivery/http/handler/quiz/test"
	quizTestRepo "csw-golang/internal/domain/repository/quiz/test"
	quizTestUsecase "csw-golang/internal/domain/usecase/quiz/test"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	db.Init()

	err := jwt.NewMiddlewareConfig()
	if err != nil {
		fmt.Println(err)
	}

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

	quizTestRepo := quizTestRepo.New(db.DB)
	quizTestUsecase := quizTestUsecase.New(quizTestRepo)
	quizTestHandler := quizTestHandler.New(quizTestUsecase)

	handler := handler.Handler{
		AuthHandler:        authHandler,
		PlanHandler:        planHandler,
		MentorHandler:      mentorHandler,
		TestimonialHandler: testimonialHandler,
		PretestHandler:     pretestHandler,
		QuizTestHandler:    quizTestHandler,
	}

	router := delivery.StartRoute(handler)
	return router
}
