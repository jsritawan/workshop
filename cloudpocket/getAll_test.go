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

func TestGetAllPockets(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	// mock.ExpectPrepare("SELECT")
	row := sqlmock.NewRows([]string{"id", "name", "budget", "balance", "is_default", "description", "currency", "account_id"}).
		AddRow(1, "test_name", 100.00, 200.00, true, "Travel", "THB", 1)
	mock.ExpectPrepare("SELECT").ExpectQuery().WillReturnRows(row)
	// WillReturnResult(sqlmock.NewResult(1, 1))

	req := httptest.NewRequest(http.MethodPost, "/cloud-pockets", strings.NewReader(``))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	h := New(db)
	c := echo.New().NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	assert.NoError(t, h.GetAll(c))

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `[{"id":1
	, "name":"test_name", "budget":100.00, "balance":200.00, "isDefault":true, "description":"Travel", "current":"THB", "accountId":1
	}]`, rec.Body.String())

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)

}
