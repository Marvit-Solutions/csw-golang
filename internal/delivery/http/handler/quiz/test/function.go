package tests

import (
	"csw-golang/internal/domain/entity/request"
	"csw-golang/internal/domain/helper/paginate"
	"csw-golang/internal/domain/helper/response"
	"csw-golang/internal/domain/helper/validator"
	"fmt"
	"net/http"
	"strconv"
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
		response.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), fmt.Errorf("failed to send param: %v", validParams))
		return
	}

	req := request.QuizParamRequest{
		UserID:    userID,
		Module:    module,
		SubModule: subModule,
		TestType:  testType,
		Limit:     int64(limit),
		Offset:    offset,
	}

	data, meta, err := th.testUsecase.GetAllTests(req)
	if err != nil {
		response.NewErrorResponse(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err)
		return
	}

	if data == nil || len(*data) == 0 {
		response.NewErrorResponse(c, http.StatusNotFound, http.StatusText(http.StatusNotFound), fmt.Errorf("test not found"))
		return
	}

	pageInt, _ := strconv.ParseInt(page, 10, 64)
	meta.Page = pageInt

	response.NewSuccessResponsePaged(c, http.StatusOK, http.StatusText(http.StatusOK), data, meta, "Success get test!")
}
