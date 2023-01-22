//go:build unit

package cloudpocket

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestTransferSuccess(t *testing.T) {

	db, mock, err := sqlmock.New()
	defer db.Close()
	assert.NoError(t, err)

	row1 := sqlmock.NewRows([]string{"balance"}).
		AddRow(2000.0)
	mock.ExpectQuery("SELECT").WithArgs(1).WillReturnRows(row1)
	mock.ExpectPrepare("UPDATE ").ExpectExec().
		WithArgs(1900.0, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	row2 := sqlmock.NewRows([]string{"balance"}).
		AddRow(0.0)
	mock.ExpectQuery("SELECT").WithArgs(2).WillReturnRows(row2)
	mock.ExpectExec("UPDATE").
		WithArgs(100.0, 2).
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := httptest.NewRequest(http.MethodPost, "/cloud-pocket/:id/transfer", strings.NewReader(`{"pocketId": 2, "amount": 100.0}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	h := New(db)
	c := echo.New().NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NoError(t, h.Transfer(c))

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"balance": 100, "pocketId": 2}`, rec.Body.String())

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)

}

func TestDecimalError(t *testing.T) {
	t.Run("Add Fund", func(t *testing.T) {
		var num1 float64 = 0.1
		var num2 float64 = 0.2
		var num3 float64 = 0.3

		res := addFund(num1, num2)
		assert.Equal(t, num3, res)
	})

	t.Run("Delete Fund", func(t *testing.T) {
		var num1 float64 = 0.3
		var num2 float64 = 0.2
		var num3 float64 = 0.1

		res := deleteFund(num1, num2)
		assert.Equal(t, num3, res)
	})

}
