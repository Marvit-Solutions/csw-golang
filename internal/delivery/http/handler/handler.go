package handler

import (
	authHandler "csw-golang/internal/delivery/http/handler/auth"
	exerciseQuestionsHandler "csw-golang/internal/delivery/http/handler/exercise/questions"
	submissionHandler "csw-golang/internal/delivery/http/handler/exercise/submission"
	exerciseTestHandler "csw-golang/internal/delivery/http/handler/exercise/test"
	mentorHandler "csw-golang/internal/delivery/http/handler/mentor"
	planHandler "csw-golang/internal/delivery/http/handler/plan"
	testimonialHandler "csw-golang/internal/delivery/http/handler/testimonial"
)

type Handler struct {
	AuthHandler              *authHandler.AuthHandler
	PlanHandler              *planHandler.PlanHandler
	MentorHandler            *mentorHandler.MentorHandler
	TestimonialHandler       *testimonialHandler.TestimonialHandler
	ExerciseQuestionsHandler *exerciseQuestionsHandler.ExerciseQuestionsHandler
	ExerciseTestHandler      *exerciseTestHandler.ExerciseTestHandler
	SubmissionHandler        *submissionHandler.SubmissionHandler
}
