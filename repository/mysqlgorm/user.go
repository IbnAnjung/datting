package mysqlgorm

import (
	"errors"

	"github.com/IbnAnjung/datting/entity/user_entity"
	"github.com/IbnAnjung/datting/utils"
	"gorm.io/gorm"
)

type userGormModel struct {
	ID            int64  `gorm:"column:id;auto_generate;<-:create"`
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	Fullname      string `gorm:"column:fullname"`
	Age           int    `gorm:"column:age"`
	Gender        string `gorm:"column:gender"`
	IsPremiumUser bool   `gorm:"column:is_premium"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) FindUserByUsername(username string) (user_entity.UserModel, error) {
	m := &userGormModel{}
	err := r.db.Table("users").Select("*").
		Where("username = ?", username).
		Find(m).Error

	if err != nil {
		return user_entity.UserModel{}, err
	}

	return user_entity.UserModel{
		ID:            m.ID,
		Username:      m.Username,
		Password:      m.Password,
		Fullname:      m.Fullname,
		Age:           m.Age,
		Gender:        m.Gender,
		IsPremiumUser: m.IsPremiumUser,
	}, nil
}

func (r UserRepository) FindUserById(id int64) (user_entity.UserModel, error) {
	m := &userGormModel{}
	err := r.db.Table("users").Select("*").
		Where("id = ?", id).
		Find(m).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = utils.DataNotFoundError
		}

		return user_entity.UserModel{}, err
	}

	return user_entity.UserModel{
		ID:            m.ID,
		Username:      m.Username,
		Password:      m.Password,
		Fullname:      m.Fullname,
		Age:           m.Age,
		Gender:        m.Gender,
		IsPremiumUser: m.IsPremiumUser,
	}, nil
}

func (r UserRepository) FindUserByIds(ids []int64) ([]user_entity.UserModel, error) {
	m := []userGormModel{}
	err := r.db.Table("users").Select("*").
		Where("id in (?)", ids).
		Find(&m).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = utils.DataNotFoundError
		}

		return []user_entity.UserModel{}, err
	}

	res := []user_entity.UserModel{}
	for _, v := range m {
		res = append(res, user_entity.UserModel{
			ID:            v.ID,
			Username:      v.Username,
			Password:      v.Password,
			Fullname:      v.Fullname,
			Age:           v.Age,
			Gender:        v.Gender,
			IsPremiumUser: v.IsPremiumUser,
		})
	}
	return res, nil
}

func (r UserRepository) CreateNewUser(entity *user_entity.UserModel) error {
	m := userGormModel{
		Username:      entity.Username,
		Password:      entity.Password,
		Fullname:      entity.Fullname,
		Age:           entity.Age,
		Gender:        entity.Gender,
		IsPremiumUser: entity.IsPremiumUser,
	}

	err := r.db.Table("users").Create(&m).Error
	entity.ID = m.ID

	return err
}

func (r UserRepository) FindUser(gender string, excldeUserIds []int64) (user_entity.UserModel, error) {
	m := &userGormModel{}
	err := r.db.Table("users").Select("*").
		Where("gender = ?", gender).
		Not(map[string]interface{}{"id": excldeUserIds}).
		First(m).Error

	if err != nil {
		return user_entity.UserModel{}, err
	}

	return user_entity.UserModel{
		ID:            m.ID,
		Username:      m.Username,
		Password:      m.Password,
		Fullname:      m.Fullname,
		Age:           m.Age,
		Gender:        m.Gender,
		IsPremiumUser: m.IsPremiumUser,
	}, nil
}
