package auth

import (
	"context"

	"github.com/IbnAnjung/datting/entity/auth_entity"
	"github.com/IbnAnjung/datting/entity/user_entity"
	"github.com/IbnAnjung/datting/utils"
)

type registerInput struct {
	Username        string `json:"username" validate:"required,alphanumunicode,min=3,max=25"`
	Password        string `json:"password" validate:"required,min=5,max=50"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password,min=5,max=50"`
	FullName        string `json:"fullname" validate:"required,min=5,max=100"`
	Age             int    `json:"age" validate:"required,numeric,min=1,max=200"`
	Gender          string `json:"gender" validate:"required,gender"`
}

func (i *registerInput) set(input auth_entity.RegisterInput) {
	i.Username = input.Username
	i.Password = input.Password
	i.ConfirmPassword = input.ConfirmPassword
	i.FullName = input.FullName
	i.Age = input.Age
	i.Gender = input.Gender
}

func (u AuthUC) Register(ctx context.Context, input auth_entity.RegisterInput) (output auth_entity.RegisterOutput, err error) {
	i := &registerInput{}
	i.set(input)

	if err = u.validator.ValidateStruct(i); err != nil {
		return
	}

	// find username exists or not
	user, err := u.userRepository.FindUserByUsername(input.Username)
	if err != nil {
		return
	}

	if user.ID != 0 {
		utils.DuplicatedDataError.Message = "username already Exists"
		err = utils.DuplicatedDataError
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
