package account

import (
	"testing"

	"github.com/IbnAnjung/datting/entity/user_entity"
	enUserMock "github.com/IbnAnjung/datting/entity/user_entity/mocks"
	"github.com/IbnAnjung/datting/utils"
	"github.com/stretchr/testify/assert"
)

func TestUpgradeWithUndefinedUser(t *testing.T) {
	userId := int64(1)
	mUser := user_entity.UserModel{}
	err := utils.DataNotFoundError

	userRepo := enUserMock.NewUserRepository(t)

	userRepo.On("FindUserById", userId).Return(mUser, err)

	uc := New(userRepo)

	assert := assert.New(t)
	res := uc.UpgradeAccount(userId)

	assert.Equal(res, err)
}

func TestUpgradeWhenUserAlreadyPremium(t *testing.T) {
	userId := int64(1)
	mUser := user_entity.UserModel{
		ID:            userId,
		IsPremiumUser: true,
	}

	err := utils.UnprocessableEntityError

	userRepo := enUserMock.NewUserRepository(t)

	userRepo.On("FindUserById", userId).Return(mUser, nil)

	uc := New(userRepo)

	assert := assert.New(t)
	res := uc.UpgradeAccount(userId)

	assert.Equal(res, err)
	assert.Equal("this account is already premium", err.Error())
}

func TestErrorWhenUpdateUser(t *testing.T) {
	userId := int64(1)
	mUser := user_entity.UserModel{
		ID:            userId,
		IsPremiumUser: false,
	}

	serverErr := utils.ServerError{}

	userRepo := enUserMock.NewUserRepository(t)

	userRepo.On("FindUserById", userId).Return(mUser, nil)

	mUser.IsPremiumUser = true
	userRepo.On("UpdateUser", &mUser).Return(serverErr)

	uc := New(userRepo)

	assert := assert.New(t)
	res := uc.UpgradeAccount(userId)

	assert.Equal(res, serverErr)
}

func TestSuccessUpdateUser(t *testing.T) {
	userId := int64(1)
	mUser := user_entity.UserModel{
		ID:            userId,
		IsPremiumUser: false,
	}

	userRepo := enUserMock.NewUserRepository(t)

	userRepo.On("FindUserById", userId).Return(mUser, nil)

	mUser.IsPremiumUser = true
	userRepo.On("UpdateUser", &mUser).Return(nil)

	uc := New(userRepo)

	assert := assert.New(t)
	res := uc.UpgradeAccount(userId)

	assert.Equal(res, nil)
}
