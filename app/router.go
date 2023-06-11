package app

import (
	"net/http"

	"github.com/IbnAnjung/datting/entity/account_entity"
	"github.com/IbnAnjung/datting/entity/auth_entity"
	"github.com/IbnAnjung/datting/entity/user_entity"
	"github.com/IbnAnjung/datting/entity/user_swap_entity"
	"github.com/IbnAnjung/datting/handler/account_handler"
	"github.com/IbnAnjung/datting/handler/auth_handler"
	"github.com/IbnAnjung/datting/handler/user_handler"
	"github.com/IbnAnjung/datting/utils"
	"github.com/gin-gonic/gin"
)

func LoadGinRouter(
	auth auth_entity.AuthUseCase,
	user user_entity.UserUseCase,
	userSwap user_swap_entity.UserSwapUseCase,
	account account_entity.AccountUseCase,
	jwt auth_entity.Jwt,
	config Config,
) *gin.Engine {

	if config.App.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

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

	userHandler := user_handler.NewUserHandler(user, userSwap)
	userH := router.Group("/user").Use(AuthMiddleware())
	{
		userH.GET("/profile", userHandler.GetUserProfile)
		userH.POST("/profile/swap", userHandler.SwapUserProfile)
	}

	accountHandler := account_handler.NewAccountHandler(account)
	accountRoute := router.Group("/account").Use(AuthMiddleware())
	{
		accountRoute.POST("/upgrade", accountHandler.Upgrade)
	}

	return router
}
