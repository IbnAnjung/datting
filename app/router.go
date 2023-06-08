package app

import (
	"net/http"

	"github.com/IbnAnjung/datting/entity/auth_entity"
	"github.com/IbnAnjung/datting/entity/user_entity"
	"github.com/IbnAnjung/datting/handler/auth_handler"
	"github.com/IbnAnjung/datting/handler/user_handler"
	"github.com/IbnAnjung/datting/utils"
	"github.com/gin-gonic/gin"
)

func LoadGinRouter(
	auth auth_entity.AuthUseCase,
	user user_entity.UserUseCase,
	jwt auth_entity.Jwt,
) *gin.Engine {
	router := gin.Default()

	//healt
	router.GET("/", func(c *gin.Context) {
		utils.SuccessResponse(c, http.StatusOK, "service is up", nil)
	})

	//auth
	authHandler := auth_handler.NewAuthHandler(auth, jwt)
	authH := router.Group("/auth")
	{
		authH.POST("/register", authHandler.Register)
		authH.POST("/login", authHandler.Login)
	}

	userHandler := user_handler.NewUserHandler(user)
	userH := router.Group("/user").Use(AuthMiddleware())
	{
		userH.GET("/profile", userHandler.GerUserProfile)
	}

	return router
}
