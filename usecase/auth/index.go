package auth

import (
	"github.com/IbnAnjung/datting/entity/auth_entity"
	"github.com/IbnAnjung/datting/entity/crypt_entity"
	"github.com/IbnAnjung/datting/entity/user_entity"
	"github.com/IbnAnjung/datting/entity/validator_entity"
)

type AuthUC struct {
	userRepository user_entity.UserRepository
	validator      validator_entity.Validator
	crypt          crypt_entity.Crypt
	jwt            auth_entity.Jwt
}

func New(
	userRepository user_entity.UserRepository,
	validator validator_entity.Validator,
	crypt crypt_entity.Crypt,
	jwt auth_entity.Jwt,
) AuthUC {
	return AuthUC{
		userRepository: userRepository,
		validator:      validator,
		crypt:          crypt,
		jwt:            jwt,
	}
}
