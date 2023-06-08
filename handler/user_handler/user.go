package user_handler

import (
	"github.com/IbnAnjung/datting/entity/user_entity"
)

type UserHandler struct {
	userUC user_entity.UserUseCase
}

func NewUserHandler(
	userUC user_entity.UserUseCase,
) UserHandler {
	return UserHandler{
		userUC: userUC,
	}
}
