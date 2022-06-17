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
		_, err := service.Create(1, 1, 1, 1, 1, 1, 1, 1)

		assert.Equal(t, err, errorConflict)
	})
}
