package tests

import (
	"csw-golang/internal/domain/entity/dto"
	"csw-golang/internal/domain/entity/request"
	"csw-golang/internal/domain/helper/paginate"
	"csw-golang/internal/domain/helper/validator"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (th *TestHandler) GetAllTests(c *gin.Context) {
	userID := "8f23e5cd-0e96-4e1d-a72b-b8f85ca083ee" // user 4

	// Define page and limit
	page := c.Query("page")
	perPage := c.Query("perpage")
	limit, offset := paginate.Paginate(page, perPage)

	// Define module and sub module
	module := c.Query("module")
	module = strings.ToUpper(module)
	subModule := c.Query("submodule")
	subModule = strings.ToUpper(subModule)
	testType := c.Query("testtype")
	testType = strings.ToUpper(testType)

	validParams := map[string]interface{}{"module": "string", "submodule": "string", "testtype": "string", "page": 0, "perpage": 0}
	if !validator.ValidateParams(c, validParams) {
		c.JSON(http.StatusNotFound, dto.Fail{
			Message: "Masukkan parameter dengan benar!",
			Code:    http.StatusBadRequest,
			Status:  http.StatusText(http.StatusBadRequest),
		})
		return
	}

	req := request.QuizParamRequest{
		UserID:    userID,
		Module:    module,
		SubModule: subModule,
		TestType:  testType,
		Page:      limit,
		Offset:    offset,
	}

	response, err := th.testUsecase.GetAllTests(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Fail{
			Code:    http.StatusInternalServerError,
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: err.Error(),
		})
		return
	}

	if response == nil || len(*response) == 0 {
		c.JSON(http.StatusNotFound, dto.Fail{
			Message: "Data not found!",
			Code:    http.StatusNotFound,
			Status:  http.StatusText(http.StatusNotFound),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Success[*[]dto.QuizResponse]{
		Message: "Success get test!",
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Data:    response,
	})
}
