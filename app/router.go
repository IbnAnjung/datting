package app

import (
	"net/http"

	"github.com/IbnAnjung/datting/entity/auth_entity"
	"github.com/IbnAnjung/datting/handler"
	"github.com/IbnAnjung/datting/utils"
	"github.com/gin-gonic/gin"
)

func LoadGinRouter(
	auth auth_entity.Auth,
	jwt auth_entity.Jwt,
) *gin.Engine {
	router := gin.Default()

	//healt
	router.GET("/", func(c *gin.Context) {
		utils.SuccessResponse(c, http.StatusOK, "service is up", nil)
	})

	//auth
	authHandler := handler.NewAuthHandler(auth, jwt)
	authH := router.Group("/auth")
	{
		authH.POST("/register", authHandler.Register)
		authH.POST("/login", authHandler.Login)
	}
	return router
}
