package auth_handler

import (
	"net/http"

	"github.com/IbnAnjung/datting/entity/auth_entity"
	"github.com/IbnAnjung/datting/handler/auth_handler/dto"
	"github.com/IbnAnjung/datting/utils"
	"github.com/gin-gonic/gin"
)

func (h AuthHandler) Login(c *gin.Context) {
	var request dto.LoginRequest
	c.ShouldBindJSON(&request)

	login, err := h.authUC.Login(c, auth_entity.LoginInput{
		Username: request.Username,
		Password: request.Password,
	})

	if err != nil {
		utils.GeneralErrorResponse(c, err, "Fail to login")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Success Login", dto.LoginResponse{
		ID:       login.ID,
		Username: login.Username,
		FullName: login.FullName,
		Age:      login.Age,
		Gender:   login.Gender,
		JwtToken: login.JwtToken,
	})
}
