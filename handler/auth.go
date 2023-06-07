package handler

import (
	"fmt"
	"net/http"

	"github.com/IbnAnjung/datting/entity/auth_entity"
	"github.com/IbnAnjung/datting/handler/request"
	"github.com/IbnAnjung/datting/handler/response"
	"github.com/IbnAnjung/datting/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUC auth_entity.AuthUseCase
	jwt    auth_entity.Jwt
}

func NewAuthHandler(
	authUC auth_entity.AuthUseCase,
	jwt auth_entity.Jwt,
) AuthHandler {
	return AuthHandler{
		authUC: authUC,
		jwt:    jwt,
	}
}

func (h AuthHandler) Register(c *gin.Context) {
	var request request.RegisterRequest
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

	jwtToken, err := h.jwt.GenerateToken(auth_entity.UserJwtClaims{
		ID:            reg.ID,
		Username:      reg.Username,
		IsPremiumUser: reg.IsPremiumUser,
	})

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Success Register new User", response.RegisterResponse{
		ID:       reg.ID,
		Username: reg.Username,
		FullName: reg.FullName,
		Age:      reg.Age,
		Gender:   reg.Gender,
		JwtToken: jwtToken,
	})
}

func (h AuthHandler) Login(c *gin.Context) {
	var request request.LoginRequest
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

	utils.SuccessResponse(c, http.StatusOK, "Success Register new User", response.LoginResponse{
		ID:       login.ID,
		Username: login.Username,
		FullName: login.FullName,
		Age:      login.Age,
		Gender:   login.Gender,
		JwtToken: login.JwtToken,
	})
}
