package testimonial

import "github.com/gin-gonic/gin"

func (testimonialHandler *TestimonialHandler) RegisterRoutes(r *gin.RouterGroup) {
	testimonialGroup := r.Group("/testimonials")
	testimonialGroup.GET("/all", testimonialHandler.GetAllTestimonials)
}
