package utils

import "github.com/gin-gonic/gin"

type response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, httpStatusCode int, message string, data interface{}) {
	c.JSON(httpStatusCode, response{
		Message: message,
		Data:    data,
	})
}
