package auth

import (
	"context"
	"errors"
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

func TestInvalidRegisterInput(t *testing.T) {
	userRepo := enUserMock.NewUserRepository(t)
	validator, _ := utils.NewValidator()
	crypt := enUtilMock.NewCrypt(t)
	jwt := enAuthMock.NewJwt(t)
	mockValidator := enUtilMock.NewValidator(t)

	uc := New(userRepo, validator, crypt, jwt)

	t.Run("test_behavior_validation_input", func(t *testing.T) {
		enInput := auth_entity.RegisterInput{
			Username:        "dealss",
			Password:        "very_secret",
			ConfirmPassword: "very_secret",
			FullName:        "Dealls",
			Age:             10,
			Gender:          "L",
		}

		ucInput := registerInput{
			Username:        enInput.Username,
			Password:        enInput.Password,
			ConfirmPassword: enInput.ConfirmPassword,
			FullName:        enInput.FullName,
			Age:             enInput.Age,
			Gender:          enInput.Gender,
		}

		mockValidator.On("ValidateStruct", &ucInput).Return(utils.ValidationError{})

		ucm := New(userRepo, mockValidator, crypt, jwt)
		ucm.Register(context.Background(), enInput)
	})

	// test username validation
	t.Run("test_user_name_required", func(t *testing.T) {
		enInput := auth_entity.RegisterInput{
			// Username:        "dealss",
			Password:        "very_secret",
			ConfirmPassword: "very_secret",
			FullName:        "Dealls",
			Age:             10,
			Gender:          "L",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_user_name_with_illegal_input", func(t *testing.T) {

		enInput := auth_entity.RegisterInput{
			Username:        "dealss國的",
			Password:        "very_secret",
			ConfirmPassword: "very_secret",
			FullName:        "Dealls",
			Age:             10,
			Gender:          "L",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_user_name_with_less_than_minimal", func(t *testing.T) {

		enInput := auth_entity.RegisterInput{
			Username:        "a",
			Password:        "very_secret",
			ConfirmPassword: "very_secret",
			FullName:        "Dealls",
			Age:             10,
			Gender:          "L",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_user_name_with_more_than_maximal", func(t *testing.T) {

		enInput := auth_entity.RegisterInput{
			Username:        "fasdklfkdsljfdsaklfaklsdfjlksjfklsadjfaklsdfjkldsajfksdvnkkdsnfksldkndfknkdvndfkvndfkvdfs",
			Password:        "very_secret",
			ConfirmPassword: "very_secret",
			FullName:        "Dealls",
			Age:             10,
			Gender:          "L",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	// end test username validation

	// test password validation
	t.Run("test_password_required", func(t *testing.T) {
		enInput := auth_entity.RegisterInput{
			Username: "dealss",
			// Password:        "very_secret",
			ConfirmPassword: "very_secret",
			FullName:        "Dealls",
			Age:             10,
			Gender:          "L",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_password_with_less_than_minimal", func(t *testing.T) {

		enInput := auth_entity.RegisterInput{
			Username:        "dealls",
			Password:        "as",
			ConfirmPassword: "very_secret",
			FullName:        "Dealls",
			Age:             10,
			Gender:          "L",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_password_with_more_than_maximal", func(t *testing.T) {

		enInput := auth_entity.RegisterInput{
			Username:        "dells",
			Password:        "asdfhjdueramksdmfk9283478u@ndklfjidf9wknkdmfvladsknkfhvjasemrfadsjfoiejrkdfvn jsnakdbfaklsdjfnvhdrghiurejnfvd,ms cvmdbsjndskm,nvs dkjfbfdnkjnmkldbgbsdfdfbdfb",
			ConfirmPassword: "very_secret",
			FullName:        "Dealls",
			Age:             10,
			Gender:          "L",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})
	// end test password validation

	// test password confirmation
	t.Run("test_confirm_password_required", func(t *testing.T) {
		enInput := auth_entity.RegisterInput{
			Username: "dealss",
			Password: "very_secret",
			// ConfirmPassword: "very_secret",
			FullName: "Dealls",
			Age:      10,
			Gender:   "L",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_confirm_password_different_by_password", func(t *testing.T) {
		enInput := auth_entity.RegisterInput{
			Username:        "dealss",
			Password:        "very_secret",
			ConfirmPassword: "wrong",
			FullName:        "Dealls",
			Age:             10,
			Gender:          "L",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	// end test password confirmation

	// test fullname
	t.Run("test_fullname_required", func(t *testing.T) {
		enInput := auth_entity.RegisterInput{
			Username:        "dealss",
			Password:        "very_secret",
			ConfirmPassword: "very_secret",
			// FullName:        "Dealls",
			Age:    10,
			Gender: "L",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_fullname_with_less_than_minimal", func(t *testing.T) {

		enInput := auth_entity.RegisterInput{
			Username:        "dealls",
			Password:        "as",
			ConfirmPassword: "very_secret",
			FullName:        "D",
			Age:             10,
			Gender:          "L",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_fullname_with_more_than_maximal", func(t *testing.T) {

		enInput := auth_entity.RegisterInput{
			Username:        "dells",
			Password:        "very_secret",
			ConfirmPassword: "very_secret",
			FullName: ` rem Ipsum is simply dummy text of the printing and
			 	typesetting industry. Lorem Ipsum has been the industry's standard dummy
			  text ever since the 1500s, when an unknown printer took a galley of
				type and scrambled it to make a type specimen book. It
				has survived not only five centuries, but also
				the leap into electronic typesetting, remaining essentially unchanged.
				It was popularised in the 1960s with the
				release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum			
			`,
			Age:    10,
			Gender: "L",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	//end test fullname

	// test age
	t.Run("test_age_required", func(t *testing.T) {
		enInput := auth_entity.RegisterInput{
			Username:        "dealss",
			Password:        "very_secret",
			ConfirmPassword: "very_secret",
			FullName:        "Dealls",
			// Age:    10,
			Gender: "L",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_age_with_less_than_minimal", func(t *testing.T) {

		enInput := auth_entity.RegisterInput{
			Username:        "dealls",
			Password:        "as",
			ConfirmPassword: "very_secret",
			FullName:        "D",
			Age:             -1,
			Gender:          "L",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_age_with_more_than_maximal", func(t *testing.T) {

		enInput := auth_entity.RegisterInput{
			Username:        "dells",
			Password:        "very_secret",
			ConfirmPassword: "very_secret",
			FullName:        "Dealls",
			Age:             300,
			Gender:          "L",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	// end test age

	// test gender
	t.Run("test_gender_required", func(t *testing.T) {
		enInput := auth_entity.RegisterInput{
			Username:        "dealss",
			Password:        "very_secret",
			ConfirmPassword: "very_secret",
			FullName:        "Dealls",
			Age:             10,
			// Gender: "L",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})

	t.Run("test_invalid_gender", func(t *testing.T) {

		enInput := auth_entity.RegisterInput{
			Username:        "dealls",
			Password:        "as",
			ConfirmPassword: "very_secret",
			FullName:        "D",
			Age:             20,
			Gender:          "Waria",
		}

		output := auth_entity.RegisterOutput{}

		res, err := uc.Register(context.Background(), enInput)

		assert.Equal(t, res, output)
		assert.EqualError(t, err, "invalid payload")
	})
	// end test gender
}

func TestRegisterWithExistUsername(t *testing.T) {

	t.Run("test_error_when_get_existing_user", func(t *testing.T) {
		enInput := auth_entity.RegisterInput{
			Username:        "dealss",
			Password:        "very_secret",
			ConfirmPassword: "very_secret",
			FullName:        "Dealls",
			Age:             10,
			Gender:          "L",
		}

		ucInput := registerInput{
			Username:        enInput.Username,
			Password:        enInput.Password,
			ConfirmPassword: enInput.ConfirmPassword,
			FullName:        enInput.FullName,
			Age:             enInput.Age,
			Gender:          enInput.Gender,
		}

		userRepo := enUserMock.NewUserRepository(t)
		validator := enUtilMock.NewValidator(t)
		crypt := enUtilMock.NewCrypt(t)
		jwt := enAuthMock.NewJwt(t)

		validator.On("ValidateStruct", &ucInput).Return(nil)
		userRepo.On("FindUserByUsername", enInput.Username).Return(user_entity.UserModel{}, errors.New("error"))

		ucm := New(userRepo, validator, crypt, jwt)
		_, err := ucm.Register(context.Background(), enInput)

		assert.NotNil(t, err)
	})

	t.Run("test_user_is_exists", func(t *testing.T) {
		enInput := auth_entity.RegisterInput{
			Username:        "dealss",
			Password:        "very_secret",
			ConfirmPassword: "very_secret",
			FullName:        "Dealls",
			Age:             10,
			Gender:          "L",
		}

		ucInput := registerInput{
			Username:        enInput.Username,
			Password:        enInput.Password,
			ConfirmPassword: enInput.ConfirmPassword,
			FullName:        enInput.FullName,
			Age:             enInput.Age,
			Gender:          enInput.Gender,
		}

		user := user_entity.UserModel{
			ID:       1,
			Username: "dealls",
		}

		userRepo := enUserMock.NewUserRepository(t)
		validator := enUtilMock.NewValidator(t)
		crypt := enUtilMock.NewCrypt(t)
		jwt := enAuthMock.NewJwt(t)

		validator.On("ValidateStruct", &ucInput).Return(nil)
		userRepo.On("FindUserByUsername", enInput.Username).Return(user, nil)

		ucm := New(userRepo, validator, crypt, jwt)
		_, err := ucm.Register(context.Background(), enInput)

		dErr := utils.DuplicatedDataError
		dErr.Message = "username already exists"

		assert.NotNil(t, err)
		assert.Equal(t, err, dErr)
	})
}

func TestValidRegister(t *testing.T) {

	t.Run("test_error_hashing", func(t *testing.T) {
		enInput := auth_entity.RegisterInput{
			Username:        "dealss",
			Password:        "very_secret",
			ConfirmPassword: "very_secret",
			FullName:        "Dealls",
			Age:             10,
			Gender:          "L",
		}

		ucInput := registerInput{
			Username:        enInput.Username,
			Password:        enInput.Password,
			ConfirmPassword: enInput.ConfirmPassword,
			FullName:        enInput.FullName,
			Age:             enInput.Age,
			Gender:          enInput.Gender,
		}

		user := user_entity.UserModel{}

		userRepo := enUserMock.NewUserRepository(t)
		validator := enUtilMock.NewValidator(t)
		crypt := enUtilMock.NewCrypt(t)
		jwt := enAuthMock.NewJwt(t)

		validator.On("ValidateStruct", &ucInput).Return(nil)
		userRepo.On("FindUserByUsername", enInput.Username).Return(user, nil)
		crypt.On("HashString", enInput.Password).Return("", errors.New("error hash"))

		ucm := New(userRepo, validator, crypt, jwt)
		_, err := ucm.Register(context.Background(), enInput)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "error hash")
	})

	t.Run("test_error_create_user", func(t *testing.T) {
		enInput := auth_entity.RegisterInput{
			Username:        "dealss",
			Password:        "very_secret",
			ConfirmPassword: "very_secret",
			FullName:        "Dealls",
			Age:             10,
			Gender:          "L",
		}

		ucInput := registerInput{
			Username:        enInput.Username,
			Password:        enInput.Password,
			ConfirmPassword: enInput.ConfirmPassword,
			FullName:        enInput.FullName,
			Age:             enInput.Age,
			Gender:          enInput.Gender,
		}

		userRepo := enUserMock.NewUserRepository(t)
		validator := enUtilMock.NewValidator(t)
		crypt := enUtilMock.NewCrypt(t)
		jwt := enAuthMock.NewJwt(t)

		validator.On("ValidateStruct", &ucInput).Return(nil)
		userRepo.On("FindUserByUsername", enInput.Username).Return(user_entity.UserModel{}, nil)
		crypt.On("HashString", enInput.Password).Return("password_hash", nil)

		// user_entity.UserModel{
		// 	Username:      "dealls",
		// 	Password:      "password_hash",
		// 	Fullname:      enInput.FullName,
		// 	Age:           enInput.Age,
		// 	Gender:        enInput.Gender,
		// 	IsPremiumUser: false,
		// }
		userRepo.On("CreateNewUser", mock.AnythingOfType("*user_entity.UserModel")).
			Return(errors.New("create user error")).
			Run(func(args mock.Arguments) {
				arg := args.Get(0).(*user_entity.UserModel)
				arg.ID = 1
			})

		ucm := New(userRepo, validator, crypt, jwt)
		_, err := ucm.Register(context.Background(), enInput)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "create user error")

	})

	t.Run("test_success_create_user", func(t *testing.T) {
		enInput := auth_entity.RegisterInput{
			Username:        "dealss",
			Password:        "very_secret",
			ConfirmPassword: "very_secret",
			FullName:        "Dealls",
			Age:             10,
			Gender:          "L",
		}

		ucInput := registerInput{
			Username:        enInput.Username,
			Password:        enInput.Password,
			ConfirmPassword: enInput.ConfirmPassword,
			FullName:        enInput.FullName,
			Age:             enInput.Age,
			Gender:          enInput.Gender,
		}

		userRepo := enUserMock.NewUserRepository(t)
		validator := enUtilMock.NewValidator(t)
		crypt := enUtilMock.NewCrypt(t)
		jwt := enAuthMock.NewJwt(t)

		validator.On("ValidateStruct", &ucInput).Return(nil)
		userRepo.On("FindUserByUsername", enInput.Username).Return(user_entity.UserModel{}, nil)
		crypt.On("HashString", enInput.Password).Return("password_hash", nil)

		user := user_entity.UserModel{
			Username:      "dealls",
			Password:      "password_hash",
			Fullname:      enInput.FullName,
			Age:           enInput.Age,
			Gender:        enInput.Gender,
			IsPremiumUser: false,
		}

		userRepo.On("CreateNewUser", mock.AnythingOfType("*user_entity.UserModel")).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(0).(*user_entity.UserModel)
				arg.ID = 1
				arg.Username = user.Username
				arg.Password = user.Password
				arg.Fullname = user.Fullname
				arg.Age = user.Age
				arg.Gender = user.Gender
				arg.IsPremiumUser = false
			})

		ucm := New(userRepo, validator, crypt, jwt)
		newUser, err := ucm.Register(context.Background(), enInput)

		assert.Nil(t, err)
		assert.IsType(t, auth_entity.RegisterOutput{}, newUser)
		assert.NotEmpty(t, newUser.ID)
	})
}
