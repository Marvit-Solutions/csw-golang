package helper

import "github.com/gin-gonic/gin"

type responseError struct {
	Message interface{} `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
}

type responsePaged struct {
	Code       int         `json:"code"`
	Status     string      `json:"status"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type responseNonPaged struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalRows  int `json:"total_rows"`
	TotalPages int `json:"total_pages"`
}

func NewErrorResponse(c *gin.Context, code int, status string, err interface{}) {
	response := new(responseError)
	response.Message = err
	response.Code = code
	response.Status = status
	c.JSON(code, &response)
}

func NewSuccessResponsePaged(c *gin.Context, code int, status string, data interface{}, pagination Pagination) {
	response := new(responsePaged)
	response.Code = code
	response.Status = status
	response.Data = data
	response.Pagination = pagination
	c.JSON(code, &response)
}
func NewSuccessResponseNonPaged(c *gin.Context, code int, status string, data interface{}) {
	response := new(responseNonPaged)
	response.Code = code
	response.Status = status
	response.Data = data
	c.JSON(code, &response)
}
