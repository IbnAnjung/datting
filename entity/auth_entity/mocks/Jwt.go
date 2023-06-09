// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Jwt is an autogenerated mock type for the Jwt type
type Jwt struct {
	mock.Mock
}

// GenerateToken provides a mock function with given fields: claims
func (_m *Jwt) GenerateToken(claims interface{}) (string, error) {
	ret := _m.Called(claims)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (string, error)); ok {
		return rf(claims)
	}
	if rf, ok := ret.Get(0).(func(interface{}) string); ok {
		r0 = rf(claims)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(claims)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseToken provides a mock function with given fields: tokenString
func (_m *Jwt) ParseToken(tokenString string) (interface{}, error) {
	ret := _m.Called(tokenString)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (interface{}, error)); ok {
		return rf(tokenString)
	}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(tokenString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(tokenString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewJwt interface {
	mock.TestingT
	Cleanup(func())
}

// NewJwt creates a new instance of Jwt. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewJwt(t mockConstructorTestingTNewJwt) *Jwt {
	mock := &Jwt{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
