package testimonial

import (
	"net/http"

	"github.com/Marvit-Solutions/csw-golang/library/helper"
	"github.com/gin-gonic/gin"
)

func (tc *testimonialHandler) GetAllTestimonials(c *gin.Context) {
	data, err := tc.testimonialUsecase.GetAllTestimonials()
	if err != nil {
		helper.NewErrorResponse(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	helper.NewSuccessResponseNonPaged(c, http.StatusOK, http.StatusText(http.StatusOK), data)
}
