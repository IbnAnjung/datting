// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	user_entity "github.com/IbnAnjung/datting/entity/user_entity"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CreateNewUser provides a mock function with given fields: _a0
func (_m *UserRepository) CreateNewUser(_a0 *user_entity.UserModel) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*user_entity.UserModel) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindUser provides a mock function with given fields: gender, excldeUserIds
func (_m *UserRepository) FindUser(gender string, excldeUserIds []int64) (user_entity.UserModel, error) {
	ret := _m.Called(gender, excldeUserIds)

	var r0 user_entity.UserModel
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []int64) (user_entity.UserModel, error)); ok {
		return rf(gender, excldeUserIds)
	}
	if rf, ok := ret.Get(0).(func(string, []int64) user_entity.UserModel); ok {
		r0 = rf(gender, excldeUserIds)
	} else {
		r0 = ret.Get(0).(user_entity.UserModel)
	}

	if rf, ok := ret.Get(1).(func(string, []int64) error); ok {
		r1 = rf(gender, excldeUserIds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserById provides a mock function with given fields: id
func (_m *UserRepository) FindUserById(id int64) (user_entity.UserModel, error) {
	ret := _m.Called(id)

	var r0 user_entity.UserModel
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (user_entity.UserModel, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) user_entity.UserModel); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(user_entity.UserModel)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByIds provides a mock function with given fields: ids
func (_m *UserRepository) FindUserByIds(ids []int64) ([]user_entity.UserModel, error) {
	ret := _m.Called(ids)

	var r0 []user_entity.UserModel
	var r1 error
	if rf, ok := ret.Get(0).(func([]int64) ([]user_entity.UserModel, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]int64) []user_entity.UserModel); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user_entity.UserModel)
		}
	}

	if rf, ok := ret.Get(1).(func([]int64) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByUsername provides a mock function with given fields: username
func (_m *UserRepository) FindUserByUsername(username string) (user_entity.UserModel, error) {
	ret := _m.Called(username)

	var r0 user_entity.UserModel
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (user_entity.UserModel, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) user_entity.UserModel); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(user_entity.UserModel)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: _a0
func (_m *UserRepository) UpdateUser(_a0 *user_entity.UserModel) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*user_entity.UserModel) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
