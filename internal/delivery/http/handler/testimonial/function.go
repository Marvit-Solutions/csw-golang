package testimonial

import (
	"csw-golang/internal/domain/helper/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (tc *TestimonialHandler) GetAllTestimonials(c *gin.Context) {
	data, err := tc.testimonialUsecase.GetAllTestimonials()
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	response.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data, "Berhasil mendapatkan testimonials!")
}
