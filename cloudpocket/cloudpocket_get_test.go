//go:build unit
// +build unit

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

func TestGetCloudpocketByID(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/cloud-pockets/:id", strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	newMockRows := sqlmock.NewRows([]string{"id", "name", "budget", "balance", "is_default", "description", "currency", "account_id"}).
		AddRow(1, "test name1", 10.9, 89.9, true, "des", "THB", 1)

	db, mock, err := sqlmock.New()

	mock.ExpectPrepare("SELECT .* FROM cloud_pockets WHERE id=\\$1").
		ExpectQuery().
		WithArgs("1").
		WillReturnRows(newMockRows)

	if err != nil {
		t.Fatalf("an error, mock expect select query '%s' was not...", err)
	}

	h := New(db)

	expected := "{\"id\":1,\"name\":\"test name1\",\"budget\":10.9,\"balance\":89.9,\"isDefault\":true,\"description\":\"des\",\"current\":\"THB\",\"accountId\":1}"

	err = h.GetCloudpocketByID(c)
	if err != nil {
		t.Fatalf("an error, act func GetCloudpocketHandler '%s' was not...", err)
	}

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
	}
}
