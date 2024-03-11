package testimonial

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/library/helper/response"

	"github.com/gin-gonic/gin"
)

func (tc *testimonialHandler) GetAllTestimonials(c *gin.Context) {
	data, err := tc.testimonialUsecase.GetAllTestimonials()
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	response.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data)
}
