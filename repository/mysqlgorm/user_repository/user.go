package user_repository

import (
	"github.com/IbnAnjung/datting/entity/user_entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) FindUserByUsername(username string) (user_entity.UserModel, error) {
	m := user_entity.UserModel{}

	return m, nil
}

func (r UserRepository) CreateNewUser(user_entity.UserModel) (user_entity.UserModel, error) {
	m := user_entity.UserModel{}

	return m, nil
}
