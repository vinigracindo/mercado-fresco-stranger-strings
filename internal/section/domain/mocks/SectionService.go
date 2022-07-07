// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/domain"
)

// SectionService is an autogenerated mock type for the SectionService type
type SectionService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId
func (_m *SectionService) Create(ctx context.Context, sectionNumber int64, currentTemperature float64, minimumTemperature float64, currentCapacity int64, minimumCapacity int64, maximumCapacity int64, warehouseId int64, productTypeId int64) (domain.SectionModel, error) {
	ret := _m.Called(ctx, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId)

	var r0 domain.SectionModel
	if rf, ok := ret.Get(0).(func(context.Context, int64, float64, float64, int64, int64, int64, int64, int64) domain.SectionModel); ok {
		r0 = rf(ctx, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId)
	} else {
		r0 = ret.Get(0).(domain.SectionModel)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, float64, float64, int64, int64, int64, int64, int64) error); ok {
		r1 = rf(ctx, sectionNumber, currentTemperature, minimumTemperature, currentCapacity, minimumCapacity, maximumCapacity, warehouseId, productTypeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *SectionService) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *SectionService) GetAll(ctx context.Context) ([]domain.SectionModel, error) {
	ret := _m.Called(ctx)

	var r0 []domain.SectionModel
	if rf, ok := ret.Get(0).(func(context.Context) []domain.SectionModel); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.SectionModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllProductCountBySection provides a mock function with given fields: ctx
func (_m *SectionService) GetAllProductCountBySection(ctx context.Context) (*[]domain.ReportProductsModel, error) {
	ret := _m.Called(ctx)

	var r0 *[]domain.ReportProductsModel
	if rf, ok := ret.Get(0).(func(context.Context) *[]domain.ReportProductsModel); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.ReportProductsModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *SectionService) GetById(ctx context.Context, id int64) (domain.SectionModel, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.SectionModel
	if rf, ok := ret.Get(0).(func(context.Context, int64) domain.SectionModel); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.SectionModel)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByIdProductCountBySection provides a mock function with given fields: ctx, id
func (_m *SectionService) GetByIdProductCountBySection(ctx context.Context, id int64) (*domain.ReportProductsModel, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.ReportProductsModel
	if rf, ok := ret.Get(0).(func(context.Context, int64) *domain.ReportProductsModel); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.ReportProductsModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCurrentCapacity provides a mock function with given fields: ctx, id, currentCapacity
func (_m *SectionService) UpdateCurrentCapacity(ctx context.Context, id int64, currentCapacity int64) (*domain.SectionModel, error) {
	ret := _m.Called(ctx, id, currentCapacity)

	var r0 *domain.SectionModel
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) *domain.SectionModel); ok {
		r0 = rf(ctx, id, currentCapacity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.SectionModel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, id, currentCapacity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewSectionService interface {
	mock.TestingT
	Cleanup(func())
}

// NewSectionService creates a new instance of SectionService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSectionService(t mockConstructorTestingTNewSectionService) *SectionService {
	mock := &SectionService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
