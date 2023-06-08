package user_entity

import "context"

type UserModel struct {
	ID            int64
	Username      string
	Password      string
	Fullname      string
	Age           int
	Gender        string
	IsPremiumUser bool
}

type Gender string

const (
	MALE   Gender = "L"
	FEMALE Gender = "P"
)

type UserRepository interface {
	FindUserByUsername(username string) (UserModel, error)
	FindUserById(id int64) (UserModel, error)
	FindUser(gender string, excldeUserIds []int64) (UserModel, error)
	CreateNewUser(*UserModel) error
}

type UserUseCase interface {
	GetRandomUserProfile(ctx context.Context, input DetailUserInput) (DetailUserOutput, error)
}
