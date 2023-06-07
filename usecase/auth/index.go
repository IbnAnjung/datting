package auth

import (
	"github.com/IbnAnjung/datting/entity/crypt_entity"
	"github.com/IbnAnjung/datting/entity/user_entity"
	"github.com/IbnAnjung/datting/entity/validator_entity"
)

type AuthUC struct {
	userRepository user_entity.UserRepository
	validator      validator_entity.Validator
	crypt          crypt_entity.Crypt
}

func New(
	userRepository user_entity.UserRepository,
	validator validator_entity.Validator,
	crypt crypt_entity.Crypt,
) AuthUC {
	return AuthUC{
		userRepository: userRepository,
		validator:      validator,
		crypt:          crypt,
	}
}
