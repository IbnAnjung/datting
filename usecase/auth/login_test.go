package auth

import (
	"context"
	"testing"

	"github.com/IbnAnjung/datting/entity/auth_entity"
	enAuthMock "github.com/IbnAnjung/datting/entity/auth_entity/mocks"
	"github.com/IbnAnjung/datting/entity/user_entity"
	enUserMock "github.com/IbnAnjung/datting/entity/user_entity/mocks"
	enUtilMock "github.com/IbnAnjung/datting/entity/util_entity/mocks"
	"github.com/IbnAnjung/datting/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInvalidInput(t *testing.T) {
	userRepo := enUserMock.NewUserRepository(t)
	validator, _ := utils.NewValidator()
	crypt := enUtilMock.NewCrypt(t)
	jwt := enAuthMock.NewJwt(t)
	mockValidator := enUtilMock.NewValidator(t)

	uc := New(userRepo, validator, crypt, jwt)

	t.Run("test_behavior_validation_input", func(t *testing.T) {
		enInput := auth_entity.LoginInput{
			Username: "dealss",
			Password: "very_secret",
		}

		ucInput := loginInput{
			Username: enInput.Username,
			Password: enInput.Password,
		}

		mockValidator.On("ValidateStruct", &ucInput).Return(utils.ValidationError{})

		ucm := New(userRepo, mockValidator, crypt, jwt)
		ucm.Login(context.Background(), enInput)
	})

	t.Run("test_user_name_required", func(t *testing.T) {
		enInput := auth_entity.LoginInput{
			Username: "",
			Password: "very_secret",
		}

		output := auth_entity.LoginOutput{}

		res, err := uc.Login(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_user_name_with_illegal_input", func(t *testing.T) {

		enInput := auth_entity.LoginInput{
			Username: "asbdj國的",
			Password: "very_secret",
		}

		output := auth_entity.LoginOutput{}

		res, err := uc.Login(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_user_name_with_less_than_minimal", func(t *testing.T) {

		enInput := auth_entity.LoginInput{
			Username: "a",
			Password: "very_secret",
		}

		output := auth_entity.LoginOutput{}

		res, err := uc.Login(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_user_name_with_more_than_maximal", func(t *testing.T) {

		enInput := auth_entity.LoginInput{
			Username: "a",
			Password: "very_secret",
		}

		output := auth_entity.LoginOutput{}

		res, err := uc.Login(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_password_required", func(t *testing.T) {
		enInput := auth_entity.LoginInput{
			Username: "dealls",
			Password: "",
		}

		output := auth_entity.LoginOutput{}

		res, err := uc.Login(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_password_with_less_than_minimal", func(t *testing.T) {

		enInput := auth_entity.LoginInput{
			Username: "dealls",
			Password: "less",
		}

		output := auth_entity.LoginOutput{}

		res, err := uc.Login(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_password_with_more_than_maximal", func(t *testing.T) {

		enInput := auth_entity.LoginInput{
			Username: "deaals",
			Password: `
			Contrary to popular belief,
			Lorem Ipsum is not simply random text. It has roots 
			in a piece of classical Latin literature from 45 BC,
			making it over 2000 years old. Richard McClintock,
			a Latin professor at Hampden-Sydney College in Virginia,
			looked up one of the more obscure Latin words, consectetur,
			from a Lorem Ipsum passage, and going through the cites of the `,
		}

		output := auth_entity.LoginOutput{}

		res, err := uc.Login(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})
}

func TestInvalidCredentials(t *testing.T) {
	t.Run("test_invalid_user_name", func(t *testing.T) {
		userRepo := enUserMock.NewUserRepository(t)
		validator := enUtilMock.NewValidator(t)
		crypt := enUtilMock.NewCrypt(t)
		jwt := enAuthMock.NewJwt(t)

		enIn := auth_entity.LoginInput{
			Username: "dealls",
			Password: "very_secret",
		}

		dErr := utils.DataNotFoundError
		validator.On("ValidateStruct", mock.Anything).Return(nil)
		userRepo.On("FindUserByUsername", enIn.Username).Return(user_entity.UserModel{}, dErr)

		uc := New(userRepo, validator, crypt, jwt)
		_, err := uc.Login(context.Background(), enIn)

		assert.ErrorIs(t, err, utils.DataNotFoundError)
		assert.EqualError(t, err, "unmatch username and password")
	})

	t.Run("", func(t *testing.T) {
		userRepo := enUserMock.NewUserRepository(t)
		validator := enUtilMock.NewValidator(t)
		crypt := enUtilMock.NewCrypt(t)
		jwt := enAuthMock.NewJwt(t)

		enIn := auth_entity.LoginInput{
			Username: "dealls",
			Password: "very_secret",
		}

		user := user_entity.UserModel{
			Password: "hash_password",
		}

		validator.On("ValidateStruct", mock.Anything).Return(nil)
		userRepo.On("FindUserByUsername", enIn.Username).Return(user, nil)
		crypt.On("VerifyHash", user.Password, enIn.Password).Return(false)

		uc := New(userRepo, validator, crypt, jwt)
		_, err := uc.Login(context.Background(), enIn)

		assert.ErrorIs(t, err, utils.DataNotFoundError)
		assert.EqualError(t, err, "unmatch username and password")
	})
}

func TestValidCredentials(t *testing.T) {
	t.Run("test_error_generate_token", func(t *testing.T) {
		userRepo := enUserMock.NewUserRepository(t)
		validator := enUtilMock.NewValidator(t)
		crypt := enUtilMock.NewCrypt(t)
		jwt := enAuthMock.NewJwt(t)

		enIn := auth_entity.LoginInput{
			Username: "dealls",
			Password: "very_secret",
		}

		user := user_entity.UserModel{
			ID:            1,
			Username:      "dealls",
			IsPremiumUser: false,
			Password:      "hash_password",
		}

		validator.On("ValidateStruct", mock.Anything).Return(nil)
		userRepo.On("FindUserByUsername", enIn.Username).Return(user, nil)
		crypt.On("VerifyHash", user.Password, enIn.Password).Return(true)
		jwt.On("GenerateToken", auth_entity.UserJwtClaims{
			ID:            user.ID,
			Username:      user.Username,
			IsPremiumUser: user.IsPremiumUser,
		}).Return("", utils.ServerError{})

		uc := New(userRepo, validator, crypt, jwt)
		_, err := uc.Login(context.Background(), enIn)

		assert.EqualError(t, err, "jwt token fail to create")
	})

	t.Run("test success login", func(t *testing.T) {
		userRepo := enUserMock.NewUserRepository(t)
		validator := enUtilMock.NewValidator(t)
		crypt := enUtilMock.NewCrypt(t)
		jwt := enAuthMock.NewJwt(t)

		enIn := auth_entity.LoginInput{
			Username: "dealls",
			Password: "very_secret",
		}

		user := user_entity.UserModel{
			ID:            1,
			Username:      "dealls",
			Fullname:      "Dealls",
			Age:           10,
			Gender:        "L",
			IsPremiumUser: false,
			Password:      "hash_password",
		}

		validator.On("ValidateStruct", mock.Anything).Return(nil)
		userRepo.On("FindUserByUsername", enIn.Username).Return(user, nil)
		crypt.On("VerifyHash", user.Password, enIn.Password).Return(true)
		jwt.On("GenerateToken", auth_entity.UserJwtClaims{
			ID:            user.ID,
			Username:      user.Username,
			IsPremiumUser: user.IsPremiumUser,
		}).Return("jwt_token_secret", nil)

		uc := New(userRepo, validator, crypt, jwt)
		output, err := uc.Login(context.Background(), enIn)

		assert.Equal(t, output, auth_entity.LoginOutput{
			ID:            user.ID,
			Username:      user.Username,
			FullName:      user.Fullname,
			Age:           user.Age,
			Gender:        user.Gender,
			IsPremiumUser: user.IsPremiumUser,
			JwtToken:      "jwt_token_secret",
		})

		assert.Equal(t, err, nil)
	})

}
