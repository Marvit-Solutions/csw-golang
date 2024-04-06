package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// ValidateRequestBody performs binding and validation of JSON request data.
// It returns an error if binding or validation fails.
func ValidateRequestBody(c *gin.Context, request interface{}) error {
	if err := c.ShouldBindJSON(request); err != nil {
		return fmt.Errorf("failed binding data: %v", err)
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return fmt.Errorf("validation error: %v", err)
	}

	return nil
}

func ValidateURLParams(c *gin.Context, params interface{}) error {
	if err := c.ShouldBindUri(params); err != nil {
		return fmt.Errorf("failed binding URI parameters: %v", err)
	}

	validate := validator.New()
	if err := validate.Struct(params); err != nil {
		return fmt.Errorf("parameter validation error: %v", err)
	}

	return nil
}
