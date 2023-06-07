package auth_handler

import (
	"fmt"
	"net/http"

	"github.com/IbnAnjung/datting/entity/auth_entity"
	"github.com/IbnAnjung/datting/handler/auth_handler/dto"
	"github.com/IbnAnjung/datting/utils"
	"github.com/gin-gonic/gin"
)

func (h AuthHandler) Login(c *gin.Context) {
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.ErrorValidationResponse(c, err.Error(), []string{})
		return
	}

	login, err := h.authUC.Login(c, auth_entity.LoginInput{
		Username: request.Username,
		Password: request.Password,
	})

	if err != nil {
		if err, ok := err.(utils.ValidationError); ok {
			validationErrors := err.Validator.GetValidationErrors()
			utils.ErrorValidationResponse(c, err.Error(), validationErrors)
			return
		}

		if err, ok := err.(utils.ClientError); ok {
			utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		fmt.Printf("user register failed %s", err.Error())
		utils.ErrorResponse(c, http.StatusInternalServerError, "user register failed")
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Success Register new User", dto.LoginResponse{
		ID:       login.ID,
		Username: login.Username,
		FullName: login.FullName,
		Age:      login.Age,
		Gender:   login.Gender,
		JwtToken: login.JwtToken,
	})
}
