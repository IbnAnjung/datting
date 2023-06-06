package app

import (
	"net/http"

	"github.com/IbnAnjung/datting/utils"
	"github.com/gin-gonic/gin"
)

func LoadGinRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		utils.SuccessResponse(c, http.StatusOK, "service is up", nil)
	})

	return router
}
