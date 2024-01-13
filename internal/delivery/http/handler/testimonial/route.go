package testimonial

import "github.com/gin-gonic/gin"

func (testimonialHandler *TestimonialHandler) RegisterRoutes(r *gin.RouterGroup) {
	authGroup := r.Group("/testimonials")
	authGroup.GET("/top", testimonialHandler.ListSixTestimonial)
	authGroup.GET("/all", testimonialHandler.GetAllTestimonials)
}
