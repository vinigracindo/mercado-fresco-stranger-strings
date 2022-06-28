package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/domain"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/domain/mocks"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/service"
)

var expectedSection = domain.SectionModel{
	SectionNumber:      "1",
	CurrentTemperature: 28.0,
	MinimumTemperature: 30.5,
	CurrentCapacity:    int64(1),
	MinimumCapacity:    int64(1),
	MaximumCapacity:    int64(1),
	WarehouseId:        int64(1),
	ProductTypeId:      int64(1),
}

var expectedUpdatedSection = domain.SectionModel{
	Id:                 int64(1),
	SectionNumber:      "1",
	CurrentTemperature: 28,
	MinimumTemperature: 30.5,
	CurrentCapacity:    int64(5),
	MinimumCapacity:    int64(1),
	MaximumCapacity:    int64(1),
	WarehouseId:        int64(1),
	ProductTypeId:      int64(1),
}

var ctx = context.Background()

func TestSectionService_Create(t *testing.T) {
	mockRepository := mocks.NewSectionRepository(t)

	t.Run("create_ok: when it contains the mandatory fields, should create a section", func(t *testing.T) {
		mockRepository.
			On("Create",
				ctx,
				expectedSection.SectionNumber,
				expectedSection.CurrentTemperature,
				expectedSection.MinimumTemperature,
				expectedSection.CurrentCapacity,
				expectedSection.MinimumCapacity,
				expectedSection.MaximumCapacity,
				expectedSection.WarehouseId,
				expectedSection.ProductTypeId,
			).
			Return(expectedSection, nil).
			Once()

		service := service.NewServiceSection(mockRepository)
		result, err := service.Create(ctx, "1", 28.0, 30.5, 1, 1, 1, 1, 1)

		assert.Nil(t, err)
		assert.Equal(t, result, expectedSection)
	})

	t.Run("create_conflict: when section_number already exists, should not create a section", func(t *testing.T) {
		errorConflict := fmt.Errorf("already a section with the code: %s", expectedSection.SectionNumber)

		mockRepository.
			On("Create",
				ctx,
				expectedSection.SectionNumber,
				expectedSection.CurrentTemperature,
				expectedSection.MinimumTemperature,
				expectedSection.CurrentCapacity,
				expectedSection.MinimumCapacity,
				expectedSection.MaximumCapacity,
				expectedSection.WarehouseId,
				expectedSection.ProductTypeId,
			).
			Return(domain.SectionModel{}, errorConflict).
			Once()

		service := service.NewServiceSection(mockRepository)
		result, err := service.Create(ctx, "1", 28.0, 30.5, 1, 1, 1, 1, 1)

		assert.Equal(t, domain.SectionModel{}, result)
		assert.Equal(t, errorConflict, err)
	})
}

func TestSectionService_GetAll(t *testing.T) {
	t.Run("get_all: when exists sections, should return a list", func(t *testing.T) {
		mockRepository := mocks.NewSectionRepository(t)
		mockRepository.
			On("GetAll", ctx).
			Return([]domain.SectionModel{expectedSection}, nil).
			Once()

		service := service.NewServiceSection(mockRepository)
		result, err := service.GetAll(ctx)

		assert.Nil(t, err)
		assert.Equal(t, []domain.SectionModel{expectedSection}, result)
	})

	t.Run("get_all_error: should return any error", func(t *testing.T) {
		mockRepository := mocks.NewSectionRepository(t)
		mockRepository.
			On("GetAll", ctx).
			Return([]domain.SectionModel{}, fmt.Errorf("any error")).
			Once()

		service := service.NewServiceSection(mockRepository)
		_, err := service.GetAll(ctx)

		assert.NotNil(t, err)
	})

}

func TestSectionService_GetById(t *testing.T) {
	mockRepository := mocks.NewSectionRepository(t)

	t.Run("find_by_id_existent: when element searched for by id exists, should return a section", func(t *testing.T) {
		mockRepository.
			On("GetById", ctx, int64(1)).
			Return(expectedSection, nil).
			Once()

		service := service.NewServiceSection(mockRepository)
		result, err := service.GetById(ctx, 1)

		assert.Nil(t, err)
		assert.Equal(t, expectedSection, result)

	})

	t.Run("find_by_id_non_existent: when the element searched for by id does not exist, should return an error", func(t *testing.T) {
		id := int64(3)
		errorNotFound := fmt.Errorf("section %d not found", id)

		mockRepository.
			On("GetById", ctx, id).
			Return(domain.SectionModel{}, errorNotFound).
			Once()

		service := service.NewServiceSection(mockRepository)
		result, err := service.GetById(ctx, id)

		assert.Equal(t, domain.SectionModel{}, result)
		assert.Equal(t, errorNotFound, err)
	})
}

func TestSectionService_Delete(t *testing.T) {
	mockRepository := mocks.NewSectionRepository(t)

	t.Run("delete_ok: when the section exist, should delete a section", func(t *testing.T) {
		mockRepository.
			On("Delete", ctx, int64(1)).
			Return(nil).
			Once()

		service := service.NewServiceSection(mockRepository)
		err := service.Delete(ctx, 1)

		assert.Nil(t, err)
	})

	t.Run("delete_non_existent: when the section does not exist, should return an error", func(t *testing.T) {
		id := int64(3)
		errorNotFound := fmt.Errorf("section %d not found", id)
		mockRepository.
			On("Delete", ctx, id).
			Return(errorNotFound).
			Once()

		service := service.NewServiceSection(mockRepository)
		err := service.Delete(ctx, id)

		assert.Equal(t, errorNotFound, err)
	})
}

func TestSectionService_Update(t *testing.T) {
	mockRepository := mocks.NewSectionRepository(t)

	t.Run("update_existent: when the data update is successful, should return the updated session", func(t *testing.T) {
		mockRepository.
			On("UpdateCurrentCapacity", ctx, int64(1), int64(5)).
			Return(expectedUpdatedSection, nil).
			Once()

		service := service.NewServiceSection(mockRepository)
		result, err := service.UpdateCurrentCapacity(ctx, int64(1), int64(5))

		assert.Nil(t, err)
		assert.Equal(t, expectedUpdatedSection.CurrentCapacity, result.CurrentCapacity)
	})

	t.Run("update_non_existent: when the element searched for by id does not exist, should return an error", func(t *testing.T) {
		id := int64(3)
		errorNotFound := fmt.Errorf("section %d not found", id)
		mockRepository.
			On("UpdateCurrentCapacity", ctx, id, int64(5)).
			Return(domain.SectionModel{}, errorNotFound).
			Once()

		service := service.NewServiceSection(mockRepository)
		result, err := service.UpdateCurrentCapacity(ctx, id, int64(5))

		assert.Equal(t, errorNotFound, err)
		assert.Equal(t, domain.SectionModel{}, result)
	})

}
