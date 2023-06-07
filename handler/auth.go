package handler

import (
	"net/http"

	"github.com/IbnAnjung/datting/entity/auth_entity"
	"github.com/IbnAnjung/datting/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUC auth_entity.Auth
}

func NewAuthHandler(
	authUC auth_entity.Auth,
) AuthHandler {
	return AuthHandler{
		authUC: authUC,
	}
}

func (h AuthHandler) Register(c *gin.Context) {
	output, _ := h.authUC.Register(c, auth_entity.RegisterInput{})

	utils.SuccessResponse(c, http.StatusOK, "Success Register new User", output)
}
