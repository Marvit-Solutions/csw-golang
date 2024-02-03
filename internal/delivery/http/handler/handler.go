package handler

import (
	authHandler "csw-golang/internal/delivery/http/handler/auth"
	mentorHandler "csw-golang/internal/delivery/http/handler/mentor"
	moduleHandler "csw-golang/internal/delivery/http/handler/modul"
	planHandler "csw-golang/internal/delivery/http/handler/plan"
	testimonialHandler "csw-golang/internal/delivery/http/handler/testimonial"
)

type Handler struct {
	AuthHandler        *authHandler.AuthHandler
	PlanHandler        *planHandler.PlanHandler
	MentorHandler      *mentorHandler.MentorHandler
	TestimonialHandler *testimonialHandler.TestimonialHandler
	ModuleHandler      *moduleHandler.ModuleHandler
}
