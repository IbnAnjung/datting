package auth

import (
	"context"
	"errors"

	"github.com/IbnAnjung/datting/entity/auth_entity"
	"github.com/IbnAnjung/datting/utils"
)

type loginInput struct {
	Username string `json:"username" validate:"required,ascii,min=3,max=25"`
	Password string `json:"password" validate:"required,min=5,max=50"`
}

func (i *loginInput) set(input auth_entity.LoginInput) {
	i.Username = input.Username
	i.Password = input.Password
}

func (u AuthUC) Login(ctx context.Context, input auth_entity.LoginInput) (output auth_entity.LoginOutput, err error) {
	i := &loginInput{}
	i.set(input)

	if err = u.validator.ValidateStruct(i); err != nil {
		return
	}

	// find username exists or not
	user, err := u.userRepository.FindUserByUsername(input.Username)
	if err != nil {
		if errors.Is(err, utils.DataNotFoundError) {
			e := utils.DataNotFoundError
			e.Message = "unmatch username and password"
			err = e
		}

		return
	}

	if ok := u.crypt.VerifyHash(user.Password, input.Password); !ok {
		e := utils.DataNotFoundError
		e.Message = "unmatch username and password"
		err = e
		return
	}

	jwtToken, err := u.jwt.GenerateToken(auth_entity.UserJwtClaims{
		ID:            user.ID,
		Username:      user.Username,
		IsPremiumUser: user.IsPremiumUser,
	})

	if err != nil {
		err = errors.New("jwt token fail to create")
		return
	}

	output = auth_entity.LoginOutput{
		ID:            user.ID,
		Username:      user.Username,
		FullName:      user.Fullname,
		Age:           user.Age,
		Gender:        user.Gender,
		IsPremiumUser: user.IsPremiumUser,
		JwtToken:      jwtToken,
	}
	return
}
