package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type validationResponse struct {
	response
	ValidationErrors []string `json:"validation_errors"`
}

func SuccessResponse(c *gin.Context, httpStatusCode int, message string, data interface{}) {
	c.JSON(httpStatusCode, response{
		Message: message,
		Data:    data,
	})
}

func ErrorValidationResponse(c *gin.Context, message string, errors []string) {
	c.JSON(http.StatusBadRequest, validationResponse{
		response: response{
			Message: message,
		},
		ValidationErrors: errors,
	})
}

func ErrorResponse(c *gin.Context, httpStatusCode int, message string) {
	c.JSON(httpStatusCode, response{
		Message: message,
	})
}
