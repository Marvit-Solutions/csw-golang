package helper

import "github.com/gin-gonic/gin"

type responseError struct {
	Message interface{} `json:"message"`
	Code    int         `json:"code"`
	Status  string      `json:"status"`
}

type responsePaged struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Meta   Meta        `json:"meta"`
}

type responseNonPaged struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type Meta struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}

func NewErrorResponse(c *gin.Context, code int, status string, err interface{}) {
	response := new(responseError)
	response.Message = err
	response.Code = code
	response.Status = status
	c.JSON(code, &response)
}

func NewSuccessResponsePaged(c *gin.Context, code int, status string, data interface{}, meta Meta) {
	response := new(responsePaged)
	response.Code = code
	response.Status = status
	response.Data = data
	response.Meta = meta
	c.JSON(code, &response)
}
func NewSuccessResponseNonPaged(c *gin.Context, code int, status string, data interface{}) {
	response := new(responseNonPaged)
	response.Code = code
	response.Status = status
	response.Data = data
	c.JSON(code, &response)
}
