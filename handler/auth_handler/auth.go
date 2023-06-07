package auth_handler

import (
	"github.com/IbnAnjung/datting/entity/auth_entity"
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
