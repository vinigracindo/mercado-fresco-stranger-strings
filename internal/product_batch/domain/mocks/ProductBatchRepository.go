// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/product_batch/domain"
)

// ProductBatchRepository is an autogenerated mock type for the ProductBatchRepository type
type ProductBatchRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, productBatch
func (_m *ProductBatchRepository) Create(ctx context.Context, productBatch *domain.ProductBatch) (*domain.ProductBatch, error) {
	ret := _m.Called(ctx, productBatch)

	var r0 *domain.ProductBatch
	if rf, ok := ret.Get(0).(func(context.Context, *domain.ProductBatch) *domain.ProductBatch); ok {
		r0 = rf(ctx, productBatch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ProductBatch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.ProductBatch) error); ok {
		r1 = rf(ctx, productBatch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductBatchRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductBatchRepository creates a new instance of ProductBatchRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductBatchRepository(t mockConstructorTestingTNewProductBatchRepository) *ProductBatchRepository {
	mock := &ProductBatchRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
