package testimonial

import (
	"github.com/Marvit-Solutions/csw-golang/internal/usecase/testimonial"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type testimonialHandler struct {
	testimonialUsecase testimonial.TestimonialUsecase
}

type TestimonialHandler interface {
	GetAllTestimonials(c *gin.Context)
}

func NewTestimonialHandler(
	db *gorm.DB,
) TestimonialHandler {
	return &testimonialHandler{
		testimonialUsecase: testimonial.NewTestimonialUsecase(db),
	}
}
