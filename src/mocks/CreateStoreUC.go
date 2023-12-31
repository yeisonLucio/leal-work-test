// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	dto "lucio.com/order-service/src/domain/dto"
)

// CreateStoreUC is an autogenerated mock type for the CreateStoreUC type
type CreateStoreUC struct {
	mock.Mock
}

// Execute provides a mock function with given fields: createStoreDTO
func (_m *CreateStoreUC) Execute(createStoreDTO dto.CreateStoreDTO) (*dto.StoreCreatedDTO, error) {
	ret := _m.Called(createStoreDTO)

	var r0 *dto.StoreCreatedDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(dto.CreateStoreDTO) (*dto.StoreCreatedDTO, error)); ok {
		return rf(createStoreDTO)
	}
	if rf, ok := ret.Get(0).(func(dto.CreateStoreDTO) *dto.StoreCreatedDTO); ok {
		r0 = rf(createStoreDTO)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.StoreCreatedDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(dto.CreateStoreDTO) error); ok {
		r1 = rf(createStoreDTO)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCreateStoreUC creates a new instance of CreateStoreUC. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCreateStoreUC(t interface {
	mock.TestingT
	Cleanup(func())
}) *CreateStoreUC {
	mock := &CreateStoreUC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
