package pretest

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/helper/debug"
	"csw-golang/internal/domain/helper/validator"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func (ph *PretestHandler) GetAllPretests(c *gin.Context) {
	// Before implement claims for user id
	userID := "8f23e5cd-0e96-4e1d-a72b-b8f85ca083ee" // user 4

	// Define page and limit
	pageParam := c.Query("page")
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	limit := 10
	offset := (page - 1) * limit

	// Remove case sensitive on query param
	module := c.Query("module")
	module = strings.ToUpper(module)

	validParams := map[string]interface{}{"module": "string", "page": 0}
	if !validator.ValidateParams(c, validParams) {
		c.JSON(http.StatusNotFound, dto.Fail{
			Message: "Masukkan parameter dengan benar!",
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	debug.TransformData(module)
	debug.TransformData(page)
	debug.TransformData(limit)
	debug.TransformData(offset)

	err, response := ph.pretestUsecase.GetAllPretests(userID, module)
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
