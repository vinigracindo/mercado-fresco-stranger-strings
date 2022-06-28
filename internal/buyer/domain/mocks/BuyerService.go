// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain"
)

// BuyerService is an autogenerated mock type for the BuyerService type
type BuyerService struct {
	mock.Mock
}

// Create provides a mock function with given fields: cardNumberId, firstName, lastName
func (_m *BuyerService) Create(cardNumberId string, firstName string, lastName string) (*domain.Buyer, error) {
	ret := _m.Called(cardNumberId, firstName, lastName)

	var r0 *domain.Buyer
	if rf, ok := ret.Get(0).(func(string, string, string) *domain.Buyer); ok {
		r0 = rf(cardNumberId, firstName, lastName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Buyer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(cardNumberId, firstName, lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *BuyerService) Delete(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *BuyerService) GetAll() ([]domain.Buyer, error) {
	ret := _m.Called()

	var r0 []domain.Buyer
	if rf, ok := ret.Get(0).(func() []domain.Buyer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Buyer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetId provides a mock function with given fields: id
func (_m *BuyerService) GetId(id int64) (*domain.Buyer, error) {
	ret := _m.Called(id)

	var r0 *domain.Buyer
	if rf, ok := ret.Get(0).(func(int64) *domain.Buyer); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Buyer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, cardNumberId, lastName
func (_m *BuyerService) Update(id int64, cardNumberId string, lastName string) (*domain.Buyer, error) {
	ret := _m.Called(id, cardNumberId, lastName)

	var r0 *domain.Buyer
	if rf, ok := ret.Get(0).(func(int64, string, string) *domain.Buyer); ok {
		r0 = rf(id, cardNumberId, lastName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Buyer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string, string) error); ok {
		r1 = rf(id, cardNumberId, lastName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBuyerService interface {
	mock.TestingT
	Cleanup(func())
}

// NewBuyerService creates a new instance of BuyerService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBuyerService(t mockConstructorTestingTNewBuyerService) *BuyerService {
	mock := &BuyerService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}