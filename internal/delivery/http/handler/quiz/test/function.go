package tests

import (
	"csw-golang/internal/delivery/http/middleware/jwt"
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
	var req request.QuizRequest

	// Define page and limit
	page := c.Query("page")
	perPage := c.Query("perpage")
	limit, offset := paginate.Paginate(page, perPage)

	// Define module, testtype and sub module
	module := c.Query("module")
	module = strings.ToUpper(module)
	subModule := c.Query("submodule")
	subModule = strings.ToUpper(subModule)
	testType := c.Query("testtype")
	testType = strings.ToUpper(testType)

	// Validate params
	validParams := map[string]interface{}{"module": "string", "submodule": "string", "testtype": "string", "page": 0, "perpage": 0}
	if !validator.ValidateParams(c, validParams) {
		response.NewErrorResponse(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), fmt.Errorf("failed to send param: %v", validParams))
		return
	}

	// Get user id from jwt token
	authenticatedUser, err := jwt.GetAuthenticatedUser(c.Request)
	if err != nil {
		response.NewErrorResponse(c, http.StatusForbidden, http.StatusText(http.StatusForbidden), err)
	}

	req = request.QuizRequest{
		UserID:    authenticatedUser,
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
