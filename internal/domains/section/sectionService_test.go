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

	t.Run("create_ok: If it contains the required fields, it will be created", func(t *testing.T) {
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

	t.Run("create_conflict: If section_number already exists it cannot be created", func(t *testing.T) {
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

func TestSectionService_GetAll(t *testing.T) {
	var listSection = []section.Section{
		{
			Id:                 int64(1),
			SectionNumber:      int64(1),
			CurrentTemperature: int64(1),
			MinimumTemperature: int64(1),
			CurrentCapacity:    int64(1),
			MinimumCapacity:    int64(1),
			MaximumCapacity:    int64(1),
			WarehouseId:        int64(1),
			ProductTypeId:      int64(1),
		},
		{
			Id:                 int64(2),
			SectionNumber:      int64(2),
			CurrentTemperature: int64(2),
			MinimumTemperature: int64(2),
			CurrentCapacity:    int64(2),
			MinimumCapacity:    int64(2),
			MaximumCapacity:    int64(2),
			WarehouseId:        int64(2),
			ProductTypeId:      int64(2),
		},
	}

	t.Run("get_all: If the list has elements, it will return an amount of the total elements", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("GetAll").Return(listSection, nil)

		service := section.NewService(repo)
		result, err := service.GetAll()

		assert.Nil(t, err)
		assert.Equal(t, listSection, result)
	})
}

func TestSectionService_GetById(t *testing.T) {
	expectedSection := section.Section{
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

	t.Run("find_by_id_existent: If the element searched for by id exists, it will be return", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("GetById", int64(1)).Return(expectedSection, nil)

		service := section.NewService(repo)
		result, err := service.GetById(1)

		assert.Nil(t, err)
		assert.Equal(t, expectedSection, result)

	})

	t.Run("find_by_id_non_existent: If the element searched for by id does not exist, return null", func(t *testing.T) {
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

func TestSectionService_Delete(t *testing.T) {
	t.Run("delete_ok: If the deletion is successful, the item will not appear in the list", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("Delete", int64(1)).Return(nil)

		service := section.NewService(repo)
		err := service.Delete(1)

		assert.Nil(t, err)
	})

	t.Run("delete_non_existent: When the section does not exist, null will be returned", func(t *testing.T) {
		id := int64(3)
		errorNotFound := fmt.Errorf("Section %d not found", id)
		repo := mocks.NewRepository(t)
		repo.On("Delete", id).Return(errorNotFound)

		service := section.NewService(repo)
		err := service.Delete(id)

		assert.Equal(t, errorNotFound, err)
	})
}

func TestSectionService_Update(t *testing.T) {
	expectedSection := section.Section{
		Id:                 int64(1),
		SectionNumber:      int64(1),
		CurrentTemperature: int64(1),
		MinimumTemperature: int64(1),
		CurrentCapacity:    int64(5),
		MinimumCapacity:    int64(1),
		MaximumCapacity:    int64(1),
		WarehouseId:        int64(1),
		ProductTypeId:      int64(1),
	}

	t.Run("update_existent: When the data update is successful, the section with the updated information will be returned", func(t *testing.T) {
		repo := mocks.NewRepository(t)
		repo.On("UpdateCurrentCapacity", int64(1), int64(5)).Return(expectedSection, nil)

		service := section.NewService(repo)
		result, err := service.UpdateCurrentCapacity(int64(1), int64(5))

		assert.Nil(t, err)
		assert.Equal(t, expectedSection.CurrentCapacity, result.CurrentCapacity)
	})

	t.Run("update_non_existent: If the section to be updated does not exist, null will be returned.", func(t *testing.T) {
		id := int64(3)
		errorNotFound := fmt.Errorf("Section %d not found", id)
		repo := mocks.NewRepository(t)
		repo.On("UpdateCurrentCapacity", id, int64(5)).Return(section.Section{}, errorNotFound)

		service := section.NewService(repo)
		result, err := service.UpdateCurrentCapacity(id, int64(5))

		assert.Equal(t, errorNotFound, err)
		assert.Equal(t, section.Section{}, result)
	})

}
