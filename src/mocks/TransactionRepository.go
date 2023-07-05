// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	entities "lucio.com/order-service/src/domain/entities"
)

// TransactionRepository is an autogenerated mock type for the TransactionRepository type
type TransactionRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: transaction
func (_m *TransactionRepository) Create(transaction entities.Transaction) (*entities.Transaction, error) {
	ret := _m.Called(transaction)

	var r0 *entities.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(entities.Transaction) (*entities.Transaction, error)); ok {
		return rf(transaction)
	}
	if rf, ok := ret.Get(0).(func(entities.Transaction) *entities.Transaction); ok {
		r0 = rf(transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(entities.Transaction) error); ok {
		r1 = rf(transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTransactionRepository creates a new instance of TransactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransactionRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransactionRepository {
	mock := &TransactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}