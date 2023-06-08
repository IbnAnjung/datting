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
	FindUserByIds(ids []int64) ([]UserModel, error)
	FindUser(gender string, excldeUserIds []int64) (UserModel, error)
	CreateNewUser(*UserModel) error
	UpdateUser(*UserModel) error
}

type UserUseCase interface {
	GetRandomUserProfile(ctx context.Context, input DetailUserInput) (DetailUserOutput, error)
}

const (
	DAILY_USER_VIEW_PROFILES string = "daily_user_view_profiles:"
	DAILY_USER_SWAP_PROFILES string = "daily_user_swap_profiles:"
)

type UserCacheRepository interface {
	SetDailyUserViewProfile(ctx context.Context, userID, viewedUserId int64) error
	GetDailyUserViewProfile(ctx context.Context, userID int64) ([]int64, error)
	SetDailyUserSwapProfile(ctx context.Context, userID, swapedUserId int64) error
	GetDailyUserSwapProfile(ctx context.Context, userID int64) ([]int64, error)
}
