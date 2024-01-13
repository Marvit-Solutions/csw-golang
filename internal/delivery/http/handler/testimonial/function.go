package testimonial

import (
	"csw-golang/internal/domain/entity/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (tc *TestimonialHandler) ListSixTestimonial(c *gin.Context) {
	err, response := tc.testimonialUsecase.GetListTopSixTestimonials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Success",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    response,
	})

}

func (tc *TestimonialHandler) GetAllTestimonials(c *gin.Context) {
	err, response := tc.testimonialUsecase.GetAllTestimonials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Success",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    response,
	})

}
