package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidationError struct {
	Error     string `json:"error" example:"Name field is required"`
	Key       string `json:"key" example:"name"`
	Condition string `json:"condition" example:"required"`
}

func ValidateRequest(c *gin.Context, request interface{}) []*ValidationError {
	var errors []*ValidationError

	if err := c.ShouldBind(request); err != nil {
		validationErrs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, validationErr := range validationErrs {
				// Map field names to human-readable names if needed
				fieldName := strings.ToLower(validationErr.Field())
				validationError := &ValidationError{
					Error:     fieldName + " field is " + validationErr.Tag(),
					Key:       fieldName,
					Condition: validationErr.Tag(),
				}
				errors = append(errors, validationError)
			}
		}
	}

	return errors
}
