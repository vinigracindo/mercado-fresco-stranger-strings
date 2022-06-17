package section_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/domains/section/mocks"
)

func TestSectionService_Create(t *testing.T) {
	expectedSection := section.Section{
		SectionNumber:      int64(1),
		CurrentTemperature: int64(1),
		MinimumTemperature: int64(1),
		CurrentCapacity:    int64(1),
		MinimumCapacity:    int64(1),
		MaximumCapacity:    int64(1),
		WarehouseId:        int64(1),
		ProductTypeId:      int64(1),
	}

	t.Run("Create ok: If it contains the required fields, it will be created", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("Create",
			expectedSection.SectionNumber,
			expectedSection.CurrentTemperature,
			expectedSection.MinimumTemperature,
			expectedSection.CurrentCapacity,
			expectedSection.MinimumCapacity,
			expectedSection.MaximumCapacity,
			expectedSection.WarehouseId,
			expectedSection.ProductTypeId,
		).Return(expectedSection, nil)

		service := section.NewService(repo)
		result, err := service.Create(1, 1, 1, 1, 1, 1, 1, 1)

		assert.Nil(t, err)
		assert.Equal(t, expectedSection.SectionNumber, result.SectionNumber)
		assert.Equal(t, expectedSection.CurrentTemperature, result.CurrentTemperature)
		assert.Equal(t, expectedSection.MinimumTemperature, result.MinimumTemperature)
		assert.Equal(t, expectedSection.CurrentCapacity, result.CurrentCapacity)
		assert.Equal(t, expectedSection.MinimumCapacity, result.MinimumCapacity)
		assert.Equal(t, expectedSection.MaximumCapacity, result.MaximumCapacity)
		assert.Equal(t, expectedSection.WarehouseId, result.WarehouseId)
		assert.Equal(t, expectedSection.ProductTypeId, result.ProductTypeId)
	})

	t.Run("Create conflict: If section_number already exists it cannot be created", func(t *testing.T) {
		errorConflict := fmt.Errorf("Already a section with the code: %d", expectedSection.SectionNumber)

		repo := mocks.NewRepository(t)
		repo.On("Create",
			expectedSection.SectionNumber,
			expectedSection.CurrentTemperature,
			expectedSection.MinimumTemperature,
			expectedSection.CurrentCapacity,
			expectedSection.MinimumCapacity,
			expectedSection.MaximumCapacity,
			expectedSection.WarehouseId,
			expectedSection.ProductTypeId,
		).Return(section.Section{}, errorConflict)

		service := section.NewService(repo)
		result, err := service.Create(1, 1, 1, 1, 1, 1, 1, 1)

		assert.Equal(t, section.Section{}, result)
		assert.Equal(t, errorConflict, err)
	})
}

func TestSectionService_Get(t *testing.T) {
	var listSection = []section.Section{}

	section01 := section.Section{
		Id:                 int64(1),
		SectionNumber:      int64(1),
		CurrentTemperature: int64(1),
		MinimumTemperature: int64(1),
		CurrentCapacity:    int64(1),
		MinimumCapacity:    int64(1),
		MaximumCapacity:    int64(1),
		WarehouseId:        int64(1),
		ProductTypeId:      int64(1),
	}

	section02 := section.Section{
		Id:                 int64(2),
		SectionNumber:      int64(2),
		CurrentTemperature: int64(2),
		MinimumTemperature: int64(2),
		CurrentCapacity:    int64(2),
		MinimumCapacity:    int64(2),
		MaximumCapacity:    int64(2),
		WarehouseId:        int64(2),
		ProductTypeId:      int64(2),
	}

	listSection = append(listSection, section01)
	listSection = append(listSection, section02)

	t.Run("Get All: If the list has elements, it will return an amount of the total elements", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("GetAll").Return(listSection, nil)

		service := section.NewService(repo)
		result, err := service.GetAll()

		assert.Nil(t, err)
		assert.Equal(t, listSection, result)
	})

	t.Run("Find by id existent: If the element searched for by id exists, it will be return", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("GetById", int64(1)).Return(section01, nil)

		service := section.NewService(repo)
		result, err := service.GetById(1)

		assert.Nil(t, err)
		assert.Equal(t, section01, result)

	})

	t.Run("Find by id non existent: If the element searched for by id does not exist, return null", func(t *testing.T) {
		id := int64(3)
		errorNotFound := fmt.Errorf("Section %d not found", id)
		repo := mocks.NewRepository(t)
		repo.On("GetById", id).Return(section.Section{}, errorNotFound)

		service := section.NewService(repo)
		result, err := service.GetById(id)

		assert.Equal(t, section.Section{}, result)
		assert.Equal(t, errorNotFound, err)
	})
}
