package pretest

import (
	"csw-golang/internal/domain/entity/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ph *PretestHandler) GetAllPretests(c *gin.Context) {
	err, response := ph.pretestUsecase.GetAllPretests()
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

func (ph *PretestHandler) GetPretestById(c *gin.Context) {
	id := c.Param("id")

	err, response := ph.pretestUsecase.GetPretestById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	if response.IDPretest == "" {
		c.JSON(http.StatusNotFound, dto.Fail{
			Message: "Pretest not found",
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
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
func (ph *PretestHandler) GetPretestReview(c *gin.Context) {
	id := c.Param("id")
	status := c.Query("status")

	validParams := map[string]bool{"status": true}
	for param := range c.Request.URL.Query() {
		if !validParams[param] {
			c.JSON(http.StatusNotFound, dto.Fail{
				Message: "Masukkan parameter dengan benar!",
				Code:    http.StatusBadRequest,
				Status:  http.StatusText(http.StatusBadRequest),
			})
			return
		}
	}

	err, response := ph.pretestUsecase.GetPretestReview(id, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	if response.IDPretest == "" {
		c.JSON(http.StatusNotFound, dto.Fail{
			Message: "Pretest not found",
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
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
