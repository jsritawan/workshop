//go:build integration

package cloudpocket

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kkgo-software-engineering/workshop/config"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestTransferIT(t *testing.T) {
	e := echo.New()

	cfg := config.New().All()
	sql, err := sql.Open("postgres", cfg.DBConnection)
	if err != nil {
		t.Error(err)
	}

	pocket := New(sql)

	e.POST("/cloud-pocket/:id/transfer", pocket.Transfer)

	reqBody := `{"amount": 100.00, "pocketId":2}`
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/cloud-pocket/%d/transfer", 1), strings.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	expected := `{"pocketId": 2, "balance": 100.0}`
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, expected, rec.Body.String())
}
