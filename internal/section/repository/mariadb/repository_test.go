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

var mockSection = domain.SectionModel{
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

var id = int64(1)

func TestSectionRepository_GetAll(t *testing.T) {

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

	t.Run("get_all_ok: should return all sections", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLGetAllSection)).WillReturnRows(rows)

		sectionRepository := repository.NewMariadbSectionRepository(db)

		result, err := sectionRepository.GetAll(context.Background())
		assert.NoError(t, err)

		assert.Equal(t, result[0].SectionNumber, int64(1))
		assert.Equal(t, result[1].SectionNumber, int64(2))
	})

	t.Run("get_all_scan_error: should return error when scan fail", func(t *testing.T) {
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

	t.Run("get_all_query_error: should return error when query fails", func(t *testing.T) {
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

	t.Run("get_by_id_ok: should return section by id", func(t *testing.T) {
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

	t.Run("get_by_id_scan_error: should return error when scan fail", func(t *testing.T) {
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

	t.Run("get_by_id_not_found: should return error when section not found", func(t *testing.T) {
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
	t.Run("create_ok: should create section", func(t *testing.T) {
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
		assert.Equal(t, result, mockSection)
	})

	t.Run("create_query_error: should return error when query execution fails", func(t *testing.T) {
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
	t.Run("delete_ok: should delete section", func(t *testing.T) {
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

	t.Run("delete_not_found: should return error when section not found", func(t *testing.T) {
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

	t.Run("delete_query_error: should return error when query execution fails", func(t *testing.T) {
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
	newCurrentCapacity := int64(1)

	t.Run("update_not_found: should return error when section not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLUpdateCurrentCapacitySection)).
			WithArgs(id, newCurrentCapacity).
			WillReturnResult(sqlmock.NewResult(0, 0))

		sectionRepository := repository.NewMariadbSectionRepository(db)

		_, err = sectionRepository.UpdateCurrentCapacity(context.Background(), &mockSection)

		assert.Error(t, err)
	})

	t.Run("update_query_error: should return error when query execution fails ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLUpdateCurrentCapacitySection)).
			WithArgs(id).
			WillReturnError(errors.New("any error"))

		sectionRepository := repository.NewMariadbSectionRepository(db)

		_, err = sectionRepository.UpdateCurrentCapacity(context.Background(), &mockSection)

		assert.Error(t, err)
	})

	t.Run("update_ok: should update section current capacity", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLUpdateCurrentCapacitySection)).
			WithArgs(id, newCurrentCapacity).
			WillReturnResult(sqlmock.NewResult(0, 1))

		sectionRepository := repository.NewMariadbSectionRepository(db)

		_, err = sectionRepository.UpdateCurrentCapacity(context.Background(), &mockSection)

		assert.NoError(t, err)
	})
}

func TestSectionRepository_GetAllProductCountBySection(t *testing.T) {

	var expectedRecordProductBySection = []domain.ReportProductsModel{
		{
			Id:            int64(1),
			SectionNumber: int64(1),
			ProductsCount: int64(200),
		},
		{
			Id:            int64(2),
			SectionNumber: int64(2),
			ProductsCount: int64(400),
		},
	}

	rows := sqlmock.NewRows([]string{
		"id",
		"sectionNumber",
		"productsCount",
	}).AddRow(
		expectedRecordProductBySection[0].Id,
		expectedRecordProductBySection[0].SectionNumber,
		expectedRecordProductBySection[0].ProductsCount,
	).AddRow(
		expectedRecordProductBySection[1].Id,
		expectedRecordProductBySection[1].SectionNumber,
		expectedRecordProductBySection[1].ProductsCount,
	)

	t.Run("get_all_product_count_by_section: should return all", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLCountProductsBySection)).WillReturnRows(rows)

		sectionRepository := repository.NewMariadbSectionRepository(db)

		result, err := sectionRepository.GetAllProductCountBySection(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, &expectedRecordProductBySection, result)
		assert.Equal(t, &expectedRecordProductBySection, result)
	})

	t.Run("get_all_product_count_by_section_scan_error: should return error when scan fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id",
			"sectionNumber",
			"productsCount",
		}).AddRow("", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLCountProductsBySection)).WillReturnRows(rows)

		sectionRepository := repository.NewMariadbSectionRepository(db)

		_, err = sectionRepository.GetAllProductCountBySection(context.Background())

		assert.Error(t, err)
	})

	t.Run("get_all_product_count_by_section_query_error: should return error when query fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLCountProductsBySection)).WillReturnError(sql.ErrNoRows)

		sectionRepository := repository.NewMariadbSectionRepository(db)

		_, err = sectionRepository.GetAllProductCountBySection(context.Background())

		assert.Error(t, err)
	})
}

func TestSectionRepository_GetByIdProductCountBySection(t *testing.T) {
	var expectedRecordProductBySection = domain.ReportProductsModel{

		Id:            int64(1),
		SectionNumber: int64(1),
		ProductsCount: int64(200),
	}

	row := sqlmock.NewRows([]string{
		"id",
		"sectionNumber",
		"productsCount",
	}).AddRow(
		expectedRecordProductBySection.Id,
		expectedRecordProductBySection.SectionNumber,
		expectedRecordProductBySection.ProductsCount,
	)

	t.Run("get_by_id_product_count_by_section_ok: should return section by id", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLCountProductsBySectionWithSectionId)).WillReturnRows(row)

		sectionRepository := repository.NewMariadbSectionRepository(db)

		result, err := sectionRepository.GetByIdProductCountBySection(context.Background(), mockSection.Id)
		assert.NoError(t, err)

		assert.Equal(t, &expectedRecordProductBySection, result)
	})

	t.Run("get_by_id_product_count_by_section_ok: should return error when scan fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		row := sqlmock.NewRows([]string{
			"id",
			"sectionNumber",
			"productsCount",
		}).AddRow("", "", "")

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLCountProductsBySectionWithSectionId)).WillReturnRows(row)

		sectionRepository := repository.NewMariadbSectionRepository(db)

		_, err = sectionRepository.GetByIdProductCountBySection(context.Background(), id)
		assert.Error(t, err)
	})

	t.Run("get_by_id_product_count_by_section_ok: should return error when section not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLCountProductsBySectionWithSectionId)).WillReturnError(sql.ErrNoRows)

		sectionRepository := repository.NewMariadbSectionRepository(db)

		_, err = sectionRepository.GetByIdProductCountBySection(context.Background(), id)
		assert.Error(t, err)

	})
}
