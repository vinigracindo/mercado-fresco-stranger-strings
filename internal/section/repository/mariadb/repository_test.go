package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/domain"
	repository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/section/repository/mariadb"
)

func TestSectionRepository_GetAll(t *testing.T) {

	t.Run("get_all_ok: ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockSections := []domain.SectionModel{
			{
				Id:                 1,
				SectionNumber:      1,
				CurrentTemperature: 1,
				MinimumTemperature: 1,
				CurrentCapacity:    1,
				MinimumCapacity:    1,
				MaximumCapacity:    1,
				WarehouseId:        1,
				ProductTypeId:      1,
			},
			{
				Id:                 2,
				SectionNumber:      2,
				CurrentTemperature: 2,
				MinimumTemperature: 2,
				CurrentCapacity:    2,
				MinimumCapacity:    2,
				MaximumCapacity:    2,
				WarehouseId:        2,
				ProductTypeId:      2,
			},
		}

		rows := sqlmock.NewRows([]string{
			"id",
			"sectionNumber",
			"currentTemperature",
			"minimumTemperature",
			"currentCapacity",
			"minimumCapacity",
			"maximumCapacity",
			"warehouseId",
			"productTypeId",
		}).AddRow(
			mockSections[0].Id,
			mockSections[0].SectionNumber,
			mockSections[0].CurrentTemperature,
			mockSections[0].MinimumTemperature,
			mockSections[0].CurrentCapacity,
			mockSections[0].MinimumCapacity,
			mockSections[0].MaximumCapacity,
			mockSections[0].WarehouseId,
			mockSections[0].ProductTypeId,
		).AddRow(
			mockSections[1].Id,
			mockSections[1].SectionNumber,
			mockSections[1].CurrentTemperature,
			mockSections[1].MinimumTemperature,
			mockSections[1].CurrentCapacity,
			mockSections[1].MinimumCapacity,
			mockSections[1].MaximumCapacity,
			mockSections[1].WarehouseId,
			mockSections[1].ProductTypeId,
		)

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLGetAllSection)).WillReturnRows(rows)

		sectionRepository := repository.NewMariadbSectionRepository(db)

		result, err := sectionRepository.GetAll(context.Background())
		assert.NoError(t, err)

		assert.Equal(t, result[0].SectionNumber, int64(1))
		assert.Equal(t, result[1].SectionNumber, int64(2))
	})

	t.Run("get_all_scan_err: ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id",
			"sectionNumber",
			"currentTemperature",
			"minimumTemperature",
			"currentCapacity",
			"minimumCapacity",
			"maximumCapacity",
			"warehouseId",
			"productTypeId",
		}).AddRow("", "", "", "", "", "", "", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLGetAllSection)).WillReturnRows(rows)

		sectionRepository := repository.NewMariadbSectionRepository(db)

		_, err = sectionRepository.GetAll(context.Background())
		assert.Error(t, err)

	})

	t.Run("get_all_select_err: ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLGetAllSection)).WillReturnError(sql.ErrNoRows)

		sectionRepository := repository.NewMariadbSectionRepository(db)

		_, err = sectionRepository.GetAll(context.Background())
		assert.Error(t, err)

	})
}

func TestSectionRepository_GetById(t *testing.T) {

	t.Run("get_by_id_ok: ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mockSection := domain.SectionModel{
			Id:                 1,
			SectionNumber:      1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseId:        1,
			ProductTypeId:      1,
		}

		row := sqlmock.NewRows([]string{
			"id",
			"sectionNumber",
			"currentTemperature",
			"minimumTemperature",
			"currentCapacity",
			"minimumCapacity",
			"maximumCapacity",
			"warehouseId",
			"productTypeId",
		}).AddRow(
			mockSection.Id,
			mockSection.SectionNumber,
			mockSection.CurrentTemperature,
			mockSection.MinimumTemperature,
			mockSection.CurrentCapacity,
			mockSection.MinimumCapacity,
			mockSection.MaximumCapacity,
			mockSection.WarehouseId,
			mockSection.ProductTypeId,
		)

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLGetByIdSection)).WillReturnRows(row)

		sectionRepository := repository.NewMariadbSectionRepository(db)

		result, err := sectionRepository.GetById(context.Background(), mockSection.Id)
		assert.NoError(t, err)

		assert.Equal(t, result.SectionNumber, int64(1))
	})

	t.Run("get_all_scan_err: ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		row := sqlmock.NewRows([]string{
			"id",
			"sectionNumber",
			"currentTemperature",
			"minimumTemperature",
			"currentCapacity",
			"minimumCapacity",
			"maximumCapacity",
			"warehouseId",
			"productTypeId",
		}).AddRow("", "", "", "", "", "", "", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLGetByIdSection)).WillReturnRows(row)

		sectionRepository := repository.NewMariadbSectionRepository(db)

		_, err = sectionRepository.GetById(context.Background(), 1)
		assert.Error(t, err)
	})

	t.Run("get_by_id_not_found: ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLGetByIdSection)).WillReturnError(sql.ErrNoRows)

		sectionRepository := repository.NewMariadbSectionRepository(db)

		_, err = sectionRepository.GetById(context.Background(), 1)
		assert.Error(t, err)

	})
}

func TestSectionRepository_Create(t *testing.T) {

	mockSection := domain.SectionModel{
		Id:                 1,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseId:        1,
		ProductTypeId:      1,
	}

	t.Run("create_ok: ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLCreateSection)).
			WithArgs(
				mockSection.SectionNumber,
				mockSection.CurrentTemperature,
				mockSection.MinimumTemperature,
				mockSection.CurrentCapacity,
				mockSection.MinimumCapacity,
				mockSection.MaximumCapacity,
				mockSection.WarehouseId,
				mockSection.ProductTypeId,
			).WillReturnResult(sqlmock.NewResult(1, 1))

		sectionRepository := repository.NewMariadbSectionRepository(db)

		result, err := sectionRepository.Create(context.Background(),
			mockSection.SectionNumber,
			mockSection.CurrentTemperature,
			mockSection.MinimumTemperature,
			mockSection.CurrentCapacity,
			mockSection.MinimumCapacity,
			mockSection.MaximumCapacity,
			mockSection.WarehouseId,
			mockSection.ProductTypeId,
		)

		assert.NoError(t, err)
		assert.Equal(t, result.SectionNumber, int64(1))
	})

	t.Run("create_fail_exec: ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLCreateSection)).
			WithArgs(0, 0, 0, 0, 0, 0, 0, 0).
			WillReturnResult(sqlmock.NewResult(1, 1))

		sectionRepository := repository.NewMariadbSectionRepository(db)
		_, err = sectionRepository.Create(context.TODO(),
			mockSection.SectionNumber,
			mockSection.CurrentTemperature,
			mockSection.MinimumTemperature,
			mockSection.CurrentCapacity,
			mockSection.MinimumCapacity,
			mockSection.MaximumCapacity,
			mockSection.WarehouseId,
			mockSection.ProductTypeId,
		)

		assert.Error(t, err)
	})

}

func TestSectionRepository_Delete(t *testing.T) {
	id := int64(1)

	t.Run("delete_ok: ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLDeleteSection)).
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(0, 1))

		sectionRepository := repository.NewMariadbSectionRepository(db)

		err = sectionRepository.Delete(context.Background(), id)

		assert.NoError(t, err)
	})

	t.Run("delete_affect_rows_0: ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLDeleteSection)).
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(0, 0))

		sectionRepository := repository.NewMariadbSectionRepository(db)

		err = sectionRepository.Delete(context.Background(), id)

		assert.Error(t, err)
	})

	t.Run("delete_not_found: ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLDeleteSection)).
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(0, 0))

		sectionRepository := repository.NewMariadbSectionRepository(db)

		err = sectionRepository.Delete(context.Background(), id)

		assert.Error(t, err)
	})

	t.Run("delete_fail: ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLDeleteSection)).
			WithArgs(id).
			WillReturnError(errors.New("any error"))

		sectionRepository := repository.NewMariadbSectionRepository(db)

		err = sectionRepository.Delete(context.Background(), id)

		assert.Error(t, err)
	})

}

func TestSectionRepository_Update(t *testing.T) {
	id := int64(1)
	newCurrentCapacity := int64(1)

	t.Run("update_not_found:", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLUpdateCurrentCapacitySection)).
			WithArgs(id, newCurrentCapacity).
			WillReturnResult(sqlmock.NewResult(0, 0))

		sectionRepository := repository.NewMariadbSectionRepository(db)

		_, err = sectionRepository.UpdateCurrentCapacity(context.Background(), id, newCurrentCapacity)

		assert.Error(t, err)
	})

	t.Run("update_fail: ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLDeleteSection)).
			WithArgs(id).
			WillReturnError(errors.New("any error"))

		sectionRepository := repository.NewMariadbSectionRepository(db)

		_, err = sectionRepository.UpdateCurrentCapacity(context.Background(), id, newCurrentCapacity)

		assert.Error(t, err)
	})

	// t.Run("update_ok:", func(t *testing.T) {
	// 	db, mock, err := sqlmock.New()
	// 	assert.NoError(t, err)
	// 	defer db.Close()

	// 	mock.ExpectQuery(regexp.QuoteMeta(repository.SQLUpdateCurrentCapacitySection)).
	// 		WithArgs(id, newCurrentCapacity).
	// 		WillReturnRows(row)

	// 	sectionRepository := repository.NewMariadbSectionRepository(db)

	// 	result, err := sectionRepository.UpdateCurrentCapacity(context.Background(), id, newCurrentCapacity)
	// 	// result, err := sectionRepository.GetById(context.Background(), id)

	// 	assert.NoError(t, err)
	// 	assert.Equal(t, result, mockSection)

	// })
}
