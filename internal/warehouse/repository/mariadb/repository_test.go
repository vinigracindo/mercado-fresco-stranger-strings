package respository

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	warehouse "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/warehouse/domain"
)

var listExpectedWarehouse []warehouse.WarehouseModel = []warehouse.WarehouseModel{
	{
		Id:                 1,
		Address:            "Avenida Teste",
		Telephone:          "31 999999999",
		WarehouseCode:      "30",
		MinimunCapacity:    10,
		MinimunTemperature: 9,
		LocalityID:         1,
	},
	{
		Id:                 2,
		Address:            "Avenida Teste 2",
		Telephone:          "31 888888888",
		WarehouseCode:      "30",
		MinimunCapacity:    77777,
		MinimunTemperature: 33333,
		LocalityID:         2,
	},
}

var expectedWarehouse warehouse.WarehouseModel = warehouse.WarehouseModel{
	Id:                 1,
	Address:            "Avenida Teste",
	Telephone:          "31 999999999",
	WarehouseCode:      "30",
	MinimunCapacity:    10,
	MinimunTemperature: 9,
	LocalityID:         1,
}

var mockWarehouse warehouse.WarehouseModel = warehouse.WarehouseModel{
	Address:            "Avenida Teste",
	Telephone:          "31 999999999",
	WarehouseCode:      "30",
	MinimunCapacity:    10,
	MinimunTemperature: 9,
	LocalityID:         1,
}

var updateWarehouse warehouse.WarehouseModel = warehouse.WarehouseModel{
	MinimunCapacity:    77777,
	MinimunTemperature: 888888,
}

func Test_repository_create(t *testing.T) {

	t.Run("create_sucess: if all the fields are correct database will create new warehouse and return it", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(CreateWarehouse)).WithArgs(
			mockWarehouse.Address,
			mockWarehouse.Telephone,
			mockWarehouse.WarehouseCode,
			mockWarehouse.MinimunCapacity,
			mockWarehouse.MinimunTemperature,
			mockWarehouse.LocalityID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		mariadbWarehouse := NewMariadbWarehouseRepository(db)

		newWarehouse, err := mariadbWarehouse.Create(context.TODO(), &mockWarehouse)

		assert.NoError(t, err)
		assert.Equal(t, expectedWarehouse, newWarehouse)
	})

	t.Run("create_error: when ExecContext return an error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(CreateWarehouse)).WithArgs(
			mockWarehouse.Address,
			mockWarehouse.Telephone,
			mockWarehouse.WarehouseCode,
			mockWarehouse.MinimunCapacity,
			mockWarehouse.MinimunTemperature,
			mockWarehouse.LocalityID,
		).WillReturnError(fmt.Errorf("some problem on the query"))

		mariadbWarehouse := NewMariadbWarehouseRepository(db)

		newWarehouse, err := mariadbWarehouse.Create(context.TODO(), &mockWarehouse)

		assert.Error(t, err)
		assert.Empty(t, newWarehouse)
	})

}

func Test_repository_update(t *testing.T) {

	t.Run("update_err_exec_query: return err because of invalid query", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		errMsg := fmt.Errorf("error: error on query")

		mock.ExpectExec(regexp.QuoteMeta(UpdateWarehouse)).WillReturnError(errMsg)

		mariadbWarehouse := NewMariadbWarehouseRepository(db)

		_, err = mariadbWarehouse.Update(context.TODO(), expectedWarehouse.Id, &updateWarehouse)

		assert.Error(t, err)

	})

	t.Run("update_success: return the entity with the fields updated", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(UpdateWarehouse)).WithArgs(
			updateWarehouse.MinimunCapacity,
			updateWarehouse.MinimunTemperature,
			expectedWarehouse.Id,
		).WillReturnResult(sqlmock.NewResult(0, 1))

		mariadbWarehouse := NewMariadbWarehouseRepository(db)

		expectReturn, err := mariadbWarehouse.Update(context.Background(), expectedWarehouse.Id, &updateWarehouse)

		assert.NoError(t, err)
		assert.Equal(t, expectReturn, updateWarehouse)
	})
}

func Test_repository_getall(t *testing.T) {
	t.Run("success_get_all: success on get all", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id",
			"adress",
			"telephone",
			"warehouse_code",
			"mininum_capacity",
			"minimum_temperature",
			"locality_id",
		}).AddRow(
			listExpectedWarehouse[0].Id,
			listExpectedWarehouse[0].Address,
			listExpectedWarehouse[0].Telephone,
			listExpectedWarehouse[0].WarehouseCode,
			listExpectedWarehouse[0].MinimunCapacity,
			listExpectedWarehouse[0].MinimunTemperature,
			listExpectedWarehouse[0].LocalityID,
		).AddRow(
			listExpectedWarehouse[1].Id,
			listExpectedWarehouse[1].Address,
			listExpectedWarehouse[1].Telephone,
			listExpectedWarehouse[1].WarehouseCode,
			listExpectedWarehouse[1].MinimunCapacity,
			listExpectedWarehouse[1].MinimunTemperature,
			listExpectedWarehouse[1].LocalityID,
		)

		mock.ExpectQuery(regexp.QuoteMeta(GetAllWarehouses)).WillReturnRows(rows)

		mariadbWarehouse := NewMariadbWarehouseRepository(db)

		expectReturn, err := mariadbWarehouse.GetAll(context.TODO())

		assert.NoError(t, err)
		assert.Equal(t, listExpectedWarehouse, expectReturn)
	})

	t.Run("error_query_get_all: return error when try exec the query", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		errMsg := fmt.Errorf("error: invalid query")

		mock.ExpectQuery(regexp.QuoteMeta(GetAllWarehouses)).WillReturnError(errMsg)

		mariadbWarehouse := NewMariadbWarehouseRepository(db)

		expectReturn, err := mariadbWarehouse.GetAll(context.TODO())

		assert.Empty(t, expectReturn)
		assert.Error(t, err)
	})

	t.Run("error_scan_get_all: return error when try to scan the query", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id",
			"adress",
			"telephone",
			"warehouse_code",
			"mininum_capacity",
			"minimum_temperature",
			"locality_id",
		}).AddRow(nil, nil, nil, nil, nil, nil, nil)

		mock.ExpectQuery(regexp.QuoteMeta(GetAllWarehouses)).WillReturnRows(rows)

		mariadbWarehouse := NewMariadbWarehouseRepository(db)

		expectReturn, err := mariadbWarehouse.GetAll(context.TODO())

		assert.Empty(t, expectReturn)
		assert.Error(t, err)
	})
}

func Test_repository_getById(t *testing.T) {
	t.Run("success_get_by_id: return the entity", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id",
			"adress",
			"telephone",
			"warehouse_code",
			"mininum_capacity",
			"minimum_temperature",
			"locality_id",
		}).AddRow(
			expectedWarehouse.Id,
			expectedWarehouse.Address,
			expectedWarehouse.Telephone,
			expectedWarehouse.WarehouseCode,
			expectedWarehouse.MinimunCapacity,
			expectedWarehouse.MinimunTemperature,
			expectedWarehouse.LocalityID,
		)

		mock.ExpectQuery(regexp.QuoteMeta(GetWarehouseById)).WithArgs(expectedWarehouse.Id).WillReturnRows(rows)

		mariadbWarehouse := NewMariadbWarehouseRepository(db)

		expectReturn, err := mariadbWarehouse.GetById(context.TODO(), expectedWarehouse.Id)

		assert.NoError(t, err)
		assert.Equal(t, expectedWarehouse, expectReturn)
	})

	t.Run("error_invalid_id: return error because of invalid id", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(GetWarehouseById)).WithArgs("aaa").WillReturnError(fmt.Errorf("aaa is not a valid id"))

		mariadbWarehouse := NewMariadbWarehouseRepository(db)

		expectReturn, err := mariadbWarehouse.GetById(context.TODO(), expectedWarehouse.Id)

		assert.Error(t, err)
		assert.Empty(t, expectReturn)
	})
}

func Test_repository_delete(t *testing.T) {
	t.Run("success_delete: return no error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(DeleteWarehouse)).WithArgs(expectedWarehouse.Id).WillReturnResult(sqlmock.NewResult(0, 1))

		mariadbWarehouse := NewMariadbWarehouseRepository(db)

		err = mariadbWarehouse.Delete(context.TODO(), expectedWarehouse.Id)

		assert.Empty(t, err)
	})

	t.Run("error_when_exec_query: return error because of an invalid query", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(DeleteWarehouse)).WithArgs(expectedWarehouse.Id).WillReturnError(fmt.Errorf("error: invalid query"))

		mariadbWarehouse := NewMariadbWarehouseRepository(db)

		err = mariadbWarehouse.Delete(context.TODO(), expectedWarehouse.Id)

		assert.Error(t, err)
	})

	t.Run("error_no_id_was_found: return error because no entity was found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(DeleteWarehouse)).WillReturnResult(sqlmock.NewResult(0, 0))

		mariadbWarehouse := NewMariadbWarehouseRepository(db)

		err = mariadbWarehouse.Delete(context.TODO(), expectedWarehouse.Id)

		assert.Error(t, err)
	})
}
