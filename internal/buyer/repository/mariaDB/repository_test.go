package repository_test

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/domain"
	repository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/buyer/repository/mariaDB"
)

var expectedBuyerList = []domain.Buyer{
	{
		Id:           1,
		CardNumberId: "1234",
		FirstName:    "First Name Teste",
		LastName:     "Last Name Teste",
	},
	{
		Id:           2,
		CardNumberId: "1235",
		FirstName:    "First Name Teste 2",
		LastName:     "Last Name Teste 2",
	},
}

var expectedBuyer = domain.Buyer{
	Id:           1,
	CardNumberId: "402323",
	FirstName:    "FirstNameTest",
	LastName:     "LastNameTest",
}

var mockBuyer = &domain.Buyer{
	Id:           1,
	CardNumberId: "402323",
	FirstName:    "FirstNameTest",
	LastName:     "LastNameTest",
}

var updateBuyer = &domain.Buyer{
	Id:           2,
	CardNumberId: "402324",
	LastName:     "LastNameTest 2",
}

var expectedPurchaseOrders = []domain.PurchaseOrdersReport{
	{
		Id:                 3,
		CardNumberId:       "40232212",
		FirstName:          "FirstNameTest",
		LastName:           "LastNameTest",
		CountBuyersRecords: 2,
	},
	{
		Id:                 4,
		CardNumberId:       "40232218",
		FirstName:          "FirstNameTest",
		LastName:           "LastNameTest",
		CountBuyersRecords: 3,
	},
}

func TestBuyerRepository_GetAll(t *testing.T) {

	rows := sqlmock.NewRows([]string{
		"id", "cardNumberId", "FirstName", "LastName",
	}).AddRow(
		expectedBuyerList[0].Id,
		expectedBuyerList[0].CardNumberId,
		expectedBuyerList[0].FirstName,
		expectedBuyerList[0].LastName,
	).AddRow(
		expectedBuyerList[1].Id,
		expectedBuyerList[1].CardNumberId,
		expectedBuyerList[1].FirstName,
		expectedBuyerList[1].LastName,
	)

	t.Run("get_all_ok: should return all buyers", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLGetAllBuyer)).WillReturnRows(rows)

		buyerRepository := repository.NewmariadbBuyerRepository(db)

		result, err := buyerRepository.GetAll(context.TODO())
		assert.NoError(t, err)

		assert.NoError(t, err)
		assert.Equal(t, &expectedBuyerList, result)

	})

	t.Run("get_all_err: should return error when query fails", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		errMsg := fmt.Errorf("error: invalid query")

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLGetAllBuyer)).WillReturnError(errMsg)

		buyerRepository := repository.NewmariadbBuyerRepository(db)

		result, err := buyerRepository.GetAll(context.TODO())

		assert.Empty(t, result)
		assert.Error(t, err)
	})

	t.Run("error: should return error when scan fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id", "cardNumberId", "FirstName", "LastName",
		}).AddRow(nil, nil, nil, nil)

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLGetAllBuyer)).WillReturnRows(rows)

		buyerRepository := repository.NewmariadbBuyerRepository(db)

		result, err := buyerRepository.GetAll(context.TODO())

		assert.Empty(t, result)
		assert.Error(t, err)

	})
}

func TestBuyerRepository_GetId(t *testing.T) {
	t.Run("getId_ok: should return buyer by id", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id", "cardNumberId", "FirstName", "LastName",
		}).AddRow(
			expectedBuyer.Id,
			expectedBuyer.CardNumberId,
			expectedBuyer.FirstName,
			expectedBuyer.LastName,
		)

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLGetByIdBuyer)).WithArgs(expectedBuyer.Id).WillReturnRows(rows)

		buyerRepository := repository.NewmariadbBuyerRepository(db)

		result, err := buyerRepository.GetId(context.TODO(), expectedBuyer.Id)

		assert.NoError(t, err)
		assert.Equal(t, &expectedBuyer, result)
	})

	t.Run("getId_error: should return error when scan fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		row := sqlmock.NewRows([]string{
			"id", "cardNumberId", "FirstName", "LastName",
		}).AddRow(nil, nil, nil, nil)

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLGetByIdBuyer)).WillReturnRows(row)

		buyerRepository := repository.NewmariadbBuyerRepository(db)

		result, err := buyerRepository.GetId(context.TODO(), expectedBuyer.Id)

		assert.Empty(t, result)
		assert.Error(t, err)

	})

	t.Run("getId_error: should return error when id not found", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLGetByIdBuyer)).WillReturnError(sql.ErrNoRows)

		buyerRepository := repository.NewmariadbBuyerRepository(db)

		result, err := buyerRepository.GetId(context.TODO(), expectedBuyer.Id)

		assert.Error(t, err)
		assert.Empty(t, result)

	})
}

func TestBuyerRepository_Create(t *testing.T) {
	t.Run("create_ok: should create buyer", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLCreateBuyer)).WithArgs(
			mockBuyer.CardNumberId,
			mockBuyer.FirstName,
			mockBuyer.LastName,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		buyerRepository := repository.NewmariadbBuyerRepository(db)

		newBuyer, err := buyerRepository.Create(context.Background(),
			mockBuyer.CardNumberId,
			mockBuyer.FirstName,
			mockBuyer.LastName,
		)

		assert.NoError(t, err)
		assert.Equal(t, newBuyer, mockBuyer)
	})

	t.Run("create_error: should return error when query execution fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLCreateBuyer)).WithArgs(
			mockBuyer.CardNumberId,
			mockBuyer.FirstName,
			mockBuyer.LastName,
		).WillReturnError(fmt.Errorf("erro"))

		buyerRepository := repository.NewmariadbBuyerRepository(db)

		newBuyer, err := buyerRepository.Create(context.Background(),
			mockBuyer.CardNumberId,
			mockBuyer.FirstName,
			mockBuyer.LastName,
		)

		assert.Error(t, err)
		assert.Empty(t, newBuyer)

	})
}

func TestBuyerRepository_Update(t *testing.T) {
	t.Run("update_ok: should update buyer", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLUpdateBuyer)).WithArgs(
			updateBuyer.CardNumberId,
			updateBuyer.LastName,
			updateBuyer.Id,
		).WillReturnResult(sqlmock.NewResult(0, 1))

		buyerRepository := repository.NewmariadbBuyerRepository(db)

		result, err := buyerRepository.Update(context.Background(),
			updateBuyer.Id,
			updateBuyer.CardNumberId,
			updateBuyer.LastName,
		)

		assert.NoError(t, err)
		assert.Equal(t, result, updateBuyer)
		fmt.Println(result)
		fmt.Println(updateBuyer)
	})

	t.Run("update_error: should return error when query execution fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		errMsg := fmt.Errorf("error: query fails")
		mock.ExpectExec(regexp.QuoteMeta(repository.SQLUpdateBuyer)).
			WillReturnError(errMsg)

		buyerRepository := repository.NewmariadbBuyerRepository(db)

		result, err := buyerRepository.Update(context.Background(), updateBuyer.Id,
			updateBuyer.CardNumberId,
			updateBuyer.LastName)

		assert.Error(t, err, result)

	})
}

func TestBuyerRepository_Delete(t *testing.T) {
	t.Run("delete_ok: should delete buyer", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLDeleteBuyer)).WithArgs(expectedBuyer.Id).WillReturnResult(sqlmock.NewResult(0, 1))

		buyerRepository := repository.NewmariadbBuyerRepository(db)

		err = buyerRepository.Delete(context.TODO(), expectedBuyer.Id)

		assert.Empty(t, err)
	})

	t.Run("delete_error: ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLDeleteBuyer)).WithArgs(expectedBuyer.Id).WillReturnError(fmt.Errorf("error"))

		buyerRepository := repository.NewmariadbBuyerRepository(db)

		err = buyerRepository.Delete(context.TODO(), expectedBuyer.Id)

		assert.Error(t, err)

	})

	t.Run("delete_error: ", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(repository.SQLDeleteBuyer)).WillReturnResult(sqlmock.NewResult(0, 0))

		buyerRepository := repository.NewmariadbBuyerRepository(db)

		err = buyerRepository.Delete(context.TODO(), expectedBuyer.Id)

		assert.Error(t, err)

	})
}

func TestBuyerRepository_GetAllPurchaseOrdersReports(t *testing.T) {
	t.Run("getAllReport_ok: should return all buyers orders", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		buyerRepository := repository.NewmariadbBuyerRepository(db)

		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id",
			"cardNumberId",
			"firstName",
			"lastName",
			"countBuyersRecords",
		})
		for _, purchaseOrderList := range expectedPurchaseOrders {
			rows = rows.AddRow(
				purchaseOrderList.Id,
				purchaseOrderList.CardNumberId,
				purchaseOrderList.FirstName,
				purchaseOrderList.LastName,
				purchaseOrderList.CountBuyersRecords,
			)
		}
		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLGetAllPurchaseOrdersReports)).
			WillReturnRows(rows)

		result, err := buyerRepository.GetAllPurchaseOrdersReports(context.TODO())

		assert.NoError(t, err)
		assert.Equal(t, &expectedPurchaseOrders, result)

	})

	t.Run("getAllReport_err: should return error when scan fails", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		buyerRepository := repository.NewmariadbBuyerRepository(db)

		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLGetAllPurchaseOrdersReports)).
			WillReturnRows(
				sqlmock.
					NewRows([]string{
						"id",
						"cardNumberId",
						"firstName",
						"lastName",
						"countBuyersRecords",
					}).AddRow(
					"", "", "", "", "",
				))

		_, err = buyerRepository.GetAllPurchaseOrdersReports(context.TODO())
		assert.Error(t, err)

	})

	t.Run("getAllReport_err: should return error when query fails", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		errMsg := fmt.Errorf("error: invalid query")

		mock.ExpectQuery(regexp.QuoteMeta(repository.SQLGetAllPurchaseOrdersReports)).
			WillReturnError(errMsg)

		buyerRepository := repository.NewmariadbBuyerRepository(db)

		_, err = buyerRepository.GetAllPurchaseOrdersReports(context.TODO())

		assert.Error(t, err)
	})
}
