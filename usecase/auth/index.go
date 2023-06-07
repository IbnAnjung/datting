package auth

import "github.com/IbnAnjung/datting/entity/user_entity"

type AuthUC struct {
	userRepository user_entity.UserRepository
}

func New(
	userRepository user_entity.UserRepository,
) AuthUC {
	return AuthUC{
		userRepository: userRepository,
	}
}
