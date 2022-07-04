package respository_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinigracindo/mercado-fresco-stranger-strings/internal/carry/domain"
	respository "github.com/vinigracindo/mercado-fresco-stranger-strings/internal/carry/repository/mariadb"
)

var mockCarry *domain.CarryModel = &domain.CarryModel{
	Id:          1,
	Cid:         1,
	CompanyName: "Mercado Livre",
	Address:     "Avenida Teste",
	Telephone:   "31 999999999",
	LocalityID:  1,
}

func Test_repository_create(t *testing.T) {
	t.Run("create_sucess: if all the fields are correct database will create new carry and return it", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(respository.QueryCreateCarry)).WithArgs(
			mockCarry.Cid,
			mockCarry.CompanyName,
			mockCarry.Address,
			mockCarry.Telephone,
			mockCarry.LocalityID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		repository := respository.NewMariadbCarryRepository(db)

		expect, err := repository.Create(context.TODO(), mockCarry)

		assert.NoError(t, err)
		assert.NotNil(t, expect)
		assert.Equal(t, expect, mockCarry)
	})

	t.Run("error_query: database return error when try to exec query", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta(respository.QueryCreateCarry)).WillReturnError(fmt.Errorf("error: invalid query"))

		repository := respository.NewMariadbCarryRepository(db)

		expect, err := repository.Create(context.TODO(), mockCarry)

		assert.Error(t, err)
		assert.Nil(t, expect)

	})

}

func Test_repository_get_by_id(t *testing.T) {
	t.Run("success_get_one: return carry from database", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{
			"id",
			"cid",
			"company_name",
			"address",
			"telephone",
			"locality_id",
		}).AddRow(
			mockCarry.Id,
			mockCarry.Cid,
			mockCarry.CompanyName,
			mockCarry.Address,
			mockCarry.Telephone,
			mockCarry.LocalityID,
		)

		mock.ExpectQuery(respository.QueryGetCarry).WithArgs(mockCarry.Id).WillReturnRows(rows)

		repository := respository.NewMariadbCarryRepository(db)

		expect, err := repository.GetById(context.TODO(), mockCarry.Id)

		assert.NoError(t, err)
		assert.NotNil(t, expect)
		assert.Equal(t, expect, mockCarry)
	})

	t.Run("error_invalid_id: return error because of invalid id or no id was found", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery(regexp.QuoteMeta(respository.QueryGetCarry)).WithArgs("BB").WillReturnError(fmt.Errorf("BB is not a valid id"))

		repository := respository.NewMariadbCarryRepository(db)

		expect, err := repository.GetById(context.TODO(), mockCarry.Id)

		assert.Error(t, err)
		assert.Nil(t, expect)
	})

}

// (cid, company_name, address, telephone, locality_id)
