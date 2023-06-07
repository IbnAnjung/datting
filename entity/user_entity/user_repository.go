package user_entity

type UserRepository interface {
	FindUserByUsername(username string) (UserModel, error)
	CreateNewUser(UserModel) (UserModel, error)
}
