// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	product "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/product"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: productCode, description, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate, productTypeId, sellerId
func (_m *ProductRepository) Create(productCode string, description string, width float64, height float64, length float64, netWeight float64, expirationRate float64, recommendedFreezingTemperature float64, freezingRate int, productTypeId int, sellerId int) (*product.Product, error) {
	ret := _m.Called(productCode, description, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate, productTypeId, sellerId)

	var r0 *product.Product
	if rf, ok := ret.Get(0).(func(string, string, float64, float64, float64, float64, float64, float64, int, int, int) *product.Product); ok {
		r0 = rf(productCode, description, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate, productTypeId, sellerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, float64, float64, float64, float64, float64, float64, int, int, int) error); ok {
		r1 = rf(productCode, description, width, height, length, netWeight, expirationRate, recommendedFreezingTemperature, freezingRate, productTypeId, sellerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *ProductRepository) Delete(id int64) error {
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
func (_m *ProductRepository) GetAll() ([]product.Product, error) {
	ret := _m.Called()

	var r0 []product.Product
	if rf, ok := ret.Get(0).(func() []product.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]product.Product)
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

// GetById provides a mock function with given fields: id
func (_m *ProductRepository) GetById(id int64) (*product.Product, error) {
	ret := _m.Called(id)

	var r0 *product.Product
	if rf, ok := ret.Get(0).(func(int64) *product.Product); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.Product)
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

// UpdateDescription provides a mock function with given fields: id, description
func (_m *ProductRepository) UpdateDescription(id int64, description string) (*product.Product, error) {
	ret := _m.Called(id, description)

	var r0 *product.Product
	if rf, ok := ret.Get(0).(func(int64, string) *product.Product); ok {
		r0 = rf(id, description)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*product.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, string) error); ok {
		r1 = rf(id, description)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductRepository creates a new instance of ProductRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductRepository(t mockConstructorTestingTNewProductRepository) *ProductRepository {
	mock := &ProductRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
