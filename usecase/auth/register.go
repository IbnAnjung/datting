package auth

import (
	"context"

	"github.com/IbnAnjung/datting/entity/auth_entity"
	"github.com/IbnAnjung/datting/entity/user_entity"
	"github.com/IbnAnjung/datting/utils"
)

func (u AuthUC) Register(ctx context.Context, input auth_entity.RegisterInput) (output auth_entity.RegisterOutput, err error) {
	if err = u.validator.ValidateStruct(input); err != nil {
		return
	}

	// find username exists or not
	user, err := u.userRepository.FindUserByUsername(input.Username)
	if err != nil {
		return
	}

	if user.ID != 0 {
		err = utils.DuplicatedDataError{}
		return
	}

	passwordHashed, err := u.crypt.HashString(input.Password)
	if err != nil {
		return
	}

	userModel := user_entity.UserModel{
		Username:      input.Username,
		Password:      passwordHashed,
		Fullname:      input.FullName,
		Age:           input.Age,
		Gender:        input.Gender,
		IsPremiumUser: false,
	}

	if err = u.userRepository.CreateNewUser(&userModel); err != nil {
		return
	}

	return auth_entity.RegisterOutput{
		ID:            userModel.ID,
		Username:      userModel.Username,
		Password:      userModel.Password,
		FullName:      userModel.Fullname,
		Age:           userModel.Age,
		Gender:        userModel.Gender,
		IsPremiumUser: userModel.IsPremiumUser,
	}, nil
}
