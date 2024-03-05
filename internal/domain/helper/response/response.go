package response

import "github.com/gin-gonic/gin"

func NewErrorResponse(c *gin.Context, code int, status string, err error) {
	response := new(ResponseError)
	response.Message = err.Error()
	response.Code = code
	response.Status = status
	c.JSON(code, &response)
}

func NewSuccessResponsePaged(c *gin.Context, code int, status string, data interface{}, meta *Meta, message string) {
	response := new(ResponsePaged[interface{}])
	response.Message = message
	response.Code = code
	response.Status = status
	response.Data = data
	response.Meta = meta
	c.JSON(code, &response)
}
func NewSuccessResponseNonPaged(c *gin.Context, code int, status string, data interface{}, message string) {
	response := new(ResponseNonPaged[interface{}])
	response.Message = message
	response.Code = code
	response.Status = status
	response.Data = data
	c.JSON(code, &response)
}
