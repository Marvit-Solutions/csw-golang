package app

import (
	db "csw-golang/pkg/database/psql"

	"csw-golang/internal/delivery"
	"csw-golang/internal/delivery/http/handler"
	authHandler "csw-golang/internal/delivery/http/handler/auth"
	exerciseQuestionsHandler "csw-golang/internal/delivery/http/handler/exercise/questions"
	exerciseSubmissionHandler "csw-golang/internal/delivery/http/handler/exercise/submission"
	exerciseTestHandler "csw-golang/internal/delivery/http/handler/exercise/test"
	mentorHandler "csw-golang/internal/delivery/http/handler/mentor"
	planHandler "csw-golang/internal/delivery/http/handler/plan"
	testimonialsHandler "csw-golang/internal/delivery/http/handler/testimonial"

	authRepo "csw-golang/internal/domain/repository/auth"
	exerciseQuestionsRepo "csw-golang/internal/domain/repository/exercise/question"
	exerciseSubmissionRepo "csw-golang/internal/domain/repository/exercise/submission"
	exerciseTestRepo "csw-golang/internal/domain/repository/exercise/test"
	mentorRepo "csw-golang/internal/domain/repository/mentor"
	planRepo "csw-golang/internal/domain/repository/plan"
	testimonialsRepo "csw-golang/internal/domain/repository/testimonial"

	authUsecase "csw-golang/internal/domain/usecase/auth"
	exerciseQuestionsUsecase "csw-golang/internal/domain/usecase/exercise/questions"
	exerciseSubmissionUsecase "csw-golang/internal/domain/usecase/exercise/submission"
	exerciseTestUsecase "csw-golang/internal/domain/usecase/exercise/test"
	mentorUsecase "csw-golang/internal/domain/usecase/mentor"
	planUsecase "csw-golang/internal/domain/usecase/plan"
	testimonialUsecase "csw-golang/internal/domain/usecase/testimonial"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	db.Init()
	authRepo := authRepo.New(db.DB)
	authUsecase := authUsecase.New(authRepo)
	authHandler := authHandler.New(authUsecase)

	exerciseQuestionsRepo := exerciseQuestionsRepo.New(db.DB)
	exerciseQuestionsUsecase := exerciseQuestionsUsecase.New(exerciseQuestionsRepo)
	exerciseQuestionsHandler := exerciseQuestionsHandler.New(exerciseQuestionsUsecase)

	exerciseSubmissionRepo := exerciseSubmissionRepo.New(db.DB)
	exerciseSubmissionUsecase := exerciseSubmissionUsecase.New(exerciseSubmissionRepo)
	exerciseSubmissionHandler := exerciseSubmissionHandler.New(exerciseSubmissionUsecase)

	exerciseTestRepo := exerciseTestRepo.New(db.DB)
	exerciseTestUsecase := exerciseTestUsecase.New(exerciseTestRepo)
	exerciseTestHandler := exerciseTestHandler.New(exerciseTestUsecase)

	planRepo := planRepo.New(db.DB)
	planUsecase := planUsecase.New(planRepo)
	planHandler := planHandler.New(planUsecase)

	mentorRepo := mentorRepo.New(db.DB)
	mentorUsecase := mentorUsecase.New(mentorRepo)
	mentorHandler := mentorHandler.New(mentorUsecase)

	testimonialRepo := testimonialsRepo.New(db.DB)
	testimonialUsecase := testimonialUsecase.New(testimonialRepo)
	testimonialHandler := testimonialsHandler.New(testimonialUsecase)

	handler := handler.Handler{
		AuthHandler:              authHandler,
		PlanHandler:              planHandler,
		MentorHandler:            mentorHandler,
		TestimonialHandler:       testimonialHandler,
		ExerciseQuestionsHandler: exerciseQuestionsHandler,
		ExerciseTestHandler:      exerciseTestHandler,
		SubmissionHandler:        exerciseSubmissionHandler,
	}

	router := delivery.StartRoute(handler)
	return router
}
