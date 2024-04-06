package helper

import "github.com/gin-gonic/gin"

type responseError struct {
	Message interface{} `json:"Message"`
	Code    int         `json:"Code"`
	Status  string      `json:"Status"`
}

type responsePaged struct {
	Code   int         `json:"Code"`
	Status string      `json:"Status"`
	Data   interface{} `json:"Data"`
	Meta   Meta        `json:"Meta"`
}

type responseNonPaged struct {
	Code   int         `json:"Code"`
	Status string      `json:"Status"`
	Data   interface{} `json:"Data"`
}

type Meta struct {
	Page  int `json:"Page"`
	Size  int `json:"Size"`
	Total int `json:"Total"`
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
