package user_swap

import (
	"github.com/IbnAnjung/datting/entity/user_entity"
	"github.com/IbnAnjung/datting/entity/user_swap_entity"
	"github.com/IbnAnjung/datting/entity/util_entity"
)

type UserSwapUC struct {
	validator           util_entity.Validator
	userRepository      user_entity.UserRepository
	userCacheRepository user_entity.UserCacheRepository
	userSwapRepository  user_swap_entity.UserSwapRepository
}

func New(
	validator util_entity.Validator,
	userRepository user_entity.UserRepository,
	userCacheRepository user_entity.UserCacheRepository,
	userSwapRepository user_swap_entity.UserSwapRepository,
) UserSwapUC {
	return UserSwapUC{
		validator:           validator,
		userRepository:      userRepository,
		userCacheRepository: userCacheRepository,
		userSwapRepository:  userSwapRepository,
	}
}
