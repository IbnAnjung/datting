package user

import (
	"github.com/IbnAnjung/datting/entity/user_entity"
	"github.com/IbnAnjung/datting/entity/util_entity"
)

type UserUC struct {
	validator           util_entity.Validator
	userRepository      user_entity.UserRepository
	userCacheRepository user_entity.UserCacheRepository
}

func New(
	validator util_entity.Validator,
	userRepository user_entity.UserRepository,
	userCacheRepository user_entity.UserCacheRepository,
) UserUC {
	return UserUC{
		validator:           validator,
		userRepository:      userRepository,
		userCacheRepository: userCacheRepository,
	}
}
