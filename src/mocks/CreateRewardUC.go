// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	dto "lucio.com/order-service/src/domain/dto"
)

// CreateRewardUC is an autogenerated mock type for the CreateRewardUC type
type CreateRewardUC struct {
	mock.Mock
}

// Execute provides a mock function with given fields: createBranchDTO
func (_m *CreateRewardUC) Execute(createBranchDTO dto.CreateRewardDTO) (*dto.RewardCreatedDTO, error) {
	ret := _m.Called(createBranchDTO)

	var r0 *dto.RewardCreatedDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(dto.CreateRewardDTO) (*dto.RewardCreatedDTO, error)); ok {
		return rf(createBranchDTO)
	}
	if rf, ok := ret.Get(0).(func(dto.CreateRewardDTO) *dto.RewardCreatedDTO); ok {
		r0 = rf(createBranchDTO)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.RewardCreatedDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(dto.CreateRewardDTO) error); ok {
		r1 = rf(createBranchDTO)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCreateRewardUC creates a new instance of CreateRewardUC. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCreateRewardUC(t interface {
	mock.TestingT
	Cleanup(func())
}) *CreateRewardUC {
	mock := &CreateRewardUC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
