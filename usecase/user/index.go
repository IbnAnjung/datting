package user

import "github.com/IbnAnjung/datting/entity/user_entity"

type UserUC struct {
	userRepository      user_entity.UserRepository
	userCacheRepository user_entity.UserCacheRepository
}

func New(
	userRepository user_entity.UserRepository,
	userCacheRepository user_entity.UserCacheRepository,
) UserUC {
	return UserUC{
		userRepository:      userRepository,
		userCacheRepository: userCacheRepository,
	}
}
