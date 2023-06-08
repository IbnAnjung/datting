package account

import (
	"github.com/IbnAnjung/datting/entity/user_entity"
)

type AccountUC struct {
	userRepository user_entity.UserRepository
}

func New(
	userRepository user_entity.UserRepository,
) AccountUC {
	return AccountUC{
		userRepository: userRepository,
	}
}
