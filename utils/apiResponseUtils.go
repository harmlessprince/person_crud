package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JSONSuccessResult struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JSONFailedValidationResult struct {
	Success bool              `json:"success" example:"false"`
	Message string            `json:"message"`
	Errors  []ValidationError `json:"errors"`
}
type JSONErrorResult struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message"`
}

// SendSuccessResponse sends a success response with a custom message.
func SendSuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": message,
		"data":    data,
	})
}

// SendErrorResponse sends an error response with a custom error message and status code.
func SendErrorResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, gin.H{
		"success": false,
		"message": message,
	})
}

// SendNotFoundErrorResponse sends an error response with a custom error message and status code.
func SendNotFoundErrorResponse(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"message": message,
	})
}

// SendValidationErrors sends validation errors as an array of error messages.
func SendValidationErrors(c *gin.Context, errors []*ValidationError) {
	c.JSON(http.StatusUnprocessableEntity, gin.H{
		"message": "Validation failed",
		"success": false,
		"errors":  errors,
	})
}
