package handler

import (
	authHandler "csw-golang/internal/delivery/http/handler/auth"
	mentorHandler "csw-golang/internal/delivery/http/handler/mentor"
	planHandler "csw-golang/internal/delivery/http/handler/plan"
	pretestHandler "csw-golang/internal/delivery/http/handler/pretest"
	quizTestHandler "csw-golang/internal/delivery/http/handler/quiz/test"
	testimonialHandler "csw-golang/internal/delivery/http/handler/testimonial"
)

type Handler struct {
	AuthHandler        *authHandler.AuthHandler
	PlanHandler        *planHandler.PlanHandler
	MentorHandler      *mentorHandler.MentorHandler
	TestimonialHandler *testimonialHandler.TestimonialHandler
	PretestHandler     *pretestHandler.PretestHandler
	QuizTestHandler    *quizTestHandler.TestHandler
}
