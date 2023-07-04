// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	dto "lucio.com/order-service/src/domain/dto"
)

// CreateCampaignUC is an autogenerated mock type for the CreateCampaignUC type
type CreateCampaignUC struct {
	mock.Mock
}

// Execute provides a mock function with given fields: createCampaignDTO
func (_m *CreateCampaignUC) Execute(createCampaignDTO dto.CreateCampaignDTO) (*dto.CampaignCreatedDTO, error) {
	ret := _m.Called(createCampaignDTO)

	var r0 *dto.CampaignCreatedDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(dto.CreateCampaignDTO) (*dto.CampaignCreatedDTO, error)); ok {
		return rf(createCampaignDTO)
	}
	if rf, ok := ret.Get(0).(func(dto.CreateCampaignDTO) *dto.CampaignCreatedDTO); ok {
		r0 = rf(createCampaignDTO)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.CampaignCreatedDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(dto.CreateCampaignDTO) error); ok {
		r1 = rf(createCampaignDTO)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCreateCampaignUC creates a new instance of CreateCampaignUC. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCreateCampaignUC(t interface {
	mock.TestingT
	Cleanup(func())
}) *CreateCampaignUC {
	mock := &CreateCampaignUC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
