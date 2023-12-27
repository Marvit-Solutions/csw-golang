package validator

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func Validation(c *gin.Context, request interface{}) error {
	if err := c.Bind(request); err != nil {
		return errors.New(fmt.Sprintf("Binding data failed: %s", err.Error()))
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			message := ""
			for _, e := range validationErrs {
				if e.Tag() == "required" {
					message = fmt.Sprintf("Masukkan %s", e.Field())
				} else if e.Tag() == "email" {
					message = "Email yang anda masukkan tidak valid"
				} else if e.Tag() == "max" && e.Field() == "Title" {
					message = "Mohon maaf, entri anda melebihi batas maksimum 65 karakter"
				} else if e.Tag() == "max" && e.Field() == "Phone" {
					message = "Nomor telepon tidak boleh lebih dari 13 digit"
				} else if e.Field() == "Phone" || e.Tag() == "min" || e.Tag() == "numeric" {
					message = fmt.Sprintf("%s tidak valid", e.Field())
				}
			}
			return errors.New(message)
		}
	}

	return nil
}
