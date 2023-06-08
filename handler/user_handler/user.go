package user_handler

import (
	"github.com/IbnAnjung/datting/entity/user_entity"
	"github.com/IbnAnjung/datting/entity/user_swap_entity"
)

type UserHandler struct {
	userUC     user_entity.UserUseCase
	userSwapUC user_swap_entity.UserSwapUseCase
}

func NewUserHandler(
	userUC user_entity.UserUseCase,
	userSwapUC user_swap_entity.UserSwapUseCase,
) UserHandler {
	return UserHandler{
		userUC:     userUC,
		userSwapUC: userSwapUC,
	}
}
