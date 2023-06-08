package auth_handler

import (
	"net/http"

	"github.com/IbnAnjung/datting/entity/auth_entity"
	"github.com/IbnAnjung/datting/handler/auth_handler/dto"
	"github.com/IbnAnjung/datting/utils"
	"github.com/gin-gonic/gin"
)

func (h AuthHandler) Register(c *gin.Context) {
	var request dto.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.ErrorValidationResponse(c, err.Error(), []string{})
		return
	}

	reg, err := h.authUC.Register(c, auth_entity.RegisterInput{
		Username:        request.Username,
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
		FullName:        request.FullName,
		Age:             request.Age,
		Gender:          request.Gender,
	})

	if err != nil {
		utils.GeneralErrorResponse(c, err, "Fail to registering new user")
		return
	}

	jwtToken, err := h.jwt.GenerateToken(auth_entity.UserJwtClaims{
		ID:            reg.ID,
		Username:      reg.Username,
		IsPremiumUser: reg.IsPremiumUser,
	})

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Success to registering new User", dto.RegisterResponse{
		ID:       reg.ID,
		Username: reg.Username,
		FullName: reg.FullName,
		Age:      reg.Age,
		Gender:   reg.Gender,
		JwtToken: jwtToken,
	})
}
