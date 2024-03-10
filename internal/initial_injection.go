package internal

import (
	"github.com/Marvit-Solutions/csw-golang/internal/app/handler/auth"
	"github.com/Marvit-Solutions/csw-golang/internal/app/handler/mentor"
	"github.com/Marvit-Solutions/csw-golang/internal/app/handler/plan"
	"github.com/Marvit-Solutions/csw-golang/internal/app/handler/testimonial"
	"github.com/Marvit-Solutions/csw-golang/library/config"

	"gorm.io/gorm"
)

type InitialInjection struct {
	Auth        auth.AuthHandler
	Mentor      mentor.MentorHandler
	Plan        plan.PlanHandler
	Testimonial testimonial.TestimonialHandler
}

// NewInitialInjection initial value dependency injection for every handler
func NewInitialInjection(sQLMaster *gorm.DB, conf config.Config) InitialInjection {
	return InitialInjection{
		Auth:        auth.NewAuthHandler(sQLMaster),
		Mentor:      mentor.NewMentorHandler(sQLMaster),
		Plan:        plan.NewPlanHandler(sQLMaster),
		Testimonial: testimonial.NewTestimonialHandler(sQLMaster),
	}
}
