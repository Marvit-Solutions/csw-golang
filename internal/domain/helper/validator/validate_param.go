package validator

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ValidateParams(c *gin.Context, validParams map[string]interface{}) bool {
	// Validasi parameter yang diperlukan
	for param, expectedType := range validParams {
		paramValue := c.Query(param)

		if paramValue == "" {
			return false
		}

		switch expectedType.(type) {
		case string:
			if _, err := strconv.Atoi(paramValue); err == nil {
				return false
			}
		case int:
			if _, err := strconv.Atoi(paramValue); err != nil {
				return false
			}
		default:
			return false
		}
	}

	for param := range c.Request.URL.Query() {
		if _, exists := validParams[param]; !exists {
			return false
		}
	}

	return true
}
