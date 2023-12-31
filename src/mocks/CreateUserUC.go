// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	dto "lucio.com/order-service/src/domain/dto"
)

// CreateUserUC is an autogenerated mock type for the CreateUserUC type
type CreateUserUC struct {
	mock.Mock
}

// Execute provides a mock function with given fields: createUserDTO
func (_m *CreateUserUC) Execute(createUserDTO dto.CreateUserDTO) (*dto.UserCreatedDTO, error) {
	ret := _m.Called(createUserDTO)

	var r0 *dto.UserCreatedDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(dto.CreateUserDTO) (*dto.UserCreatedDTO, error)); ok {
		return rf(createUserDTO)
	}
	if rf, ok := ret.Get(0).(func(dto.CreateUserDTO) *dto.UserCreatedDTO); ok {
		r0 = rf(createUserDTO)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UserCreatedDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(dto.CreateUserDTO) error); ok {
		r1 = rf(createUserDTO)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCreateUserUC creates a new instance of CreateUserUC. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCreateUserUC(t interface {
	mock.TestingT
	Cleanup(func())
}) *CreateUserUC {
	mock := &CreateUserUC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
