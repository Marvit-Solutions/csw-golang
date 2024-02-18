package pretest

import (
	"csw-golang/internal/domain/entity/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

	c.JSON(http.StatusOK, dto.Success[[]dto.GetAllPretestResponse]{
		Message: "Berhasil mendapatkan pretest!",
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

	c.JSON(http.StatusOK, dto.Success[dto.Pretest]{
		Message: "Berhasil mendapatkan pretest!",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    response,
	})
}
func (ph *PretestHandler) GetPretestReview(c *gin.Context) {
	id := c.Param("id")
	status := c.Query("status")

	// Remove case sensitive on query param
	titleCase := cases.Title(language.English)
	status = titleCase.String(status)

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

	c.JSON(http.StatusOK, dto.Success[dto.PretestReviewResponse]{
		Message: "Berhasil mendapatkan review pretest!",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    response,
	})

}

func (ph *PretestHandler) SubmitPretest(c *gin.Context) {
	var req dto.PretestSubmitRequest

	id := c.Param("id")

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	err := ph.pretestUsecase.SubmitPretest(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
			Status:  http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	err = ph.pretestUsecase.GradingPretest(id, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
			Status:  http.StatusText(http.StatusInternalServerError),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[interface{}]{
		Message: "Submit sukses!",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
	})
}
