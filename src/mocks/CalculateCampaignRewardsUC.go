// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// CalculateCampaignRewardsUC is an autogenerated mock type for the CalculateCampaignRewardsUC type
type CalculateCampaignRewardsUC struct {
	mock.Mock
}

// Execute provides a mock function with given fields: brachID, storePoints, storeCoins
func (_m *CalculateCampaignRewardsUC) Execute(brachID uint, storePoints uint, storeCoins uint) (uint, uint) {
	ret := _m.Called(brachID, storePoints, storeCoins)

	var r0 uint
	var r1 uint
	if rf, ok := ret.Get(0).(func(uint, uint, uint) (uint, uint)); ok {
		return rf(brachID, storePoints, storeCoins)
	}
	if rf, ok := ret.Get(0).(func(uint, uint, uint) uint); ok {
		r0 = rf(brachID, storePoints, storeCoins)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(uint, uint, uint) uint); ok {
		r1 = rf(brachID, storePoints, storeCoins)
	} else {
		r1 = ret.Get(1).(uint)
	}

	return r0, r1
}

// NewCalculateCampaignRewardsUC creates a new instance of CalculateCampaignRewardsUC. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCalculateCampaignRewardsUC(t interface {
	mock.TestingT
	Cleanup(func())
}) *CalculateCampaignRewardsUC {
	mock := &CalculateCampaignRewardsUC{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
