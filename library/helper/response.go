package helper

import "github.com/gin-gonic/gin"

type ResponseError struct {
	Message string `json:"Message"`
	Code    int    `json:"Code"`
	Status  string `json:"Status"`
}

type ResponsePaged[Data interface{}] struct {
	Code   int    `json:"Code"`
	Status string `json:"Status"`
	Data   Data   `json:"Data"`
	Meta   Meta   `json:"Meta"`
	// "Meta": {
	// 	"Page": "4"
	// 	"PerPage": "10",
	// 	"LastPage": "25",
	// 	"TotalPage": "50",
	//}
}

type ResponseNonPaged[Data interface{}] struct {
	Code   int    `json:"Code"`
	Status string `json:"Status"`
	Data   Data   `json:"Data"`
}

type Meta struct {
	Page      int64 `json:"Page"`
	PerPage   int64 `json:"PerPage"`
	LastPage  int64 `json:"LastPage"`
	TotalPage int64 `json:"TotalPage"`
}

func NewErrorResponse(c *gin.Context, code int, status string, err error) {
	response := new(ResponseError)
	response.Message = err.Error()
	response.Code = code
	response.Status = status
	c.JSON(code, &response)
}

func NewSuccessResponsePaged(c *gin.Context, code int, status string, data interface{}, meta Meta) {
	response := new(ResponsePaged[interface{}])
	response.Code = code
	response.Status = status
	response.Data = data
	response.Meta = meta
	c.JSON(code, &response)
}
func NewSuccessResponseNonPaged(c *gin.Context, code int, status string, data interface{}) {
	response := new(ResponseNonPaged[interface{}])
	response.Code = code
	response.Status = status
	response.Data = data
	c.JSON(code, &response)
}
