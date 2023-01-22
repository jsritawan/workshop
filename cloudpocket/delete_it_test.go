//go:build integration

package cloudpocket

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kkgo-software-engineering/workshop/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCloudPocketSuccess(t *testing.T) {
	e := echo.New()

	cfg := config.New().All()
	db, err := sql.Open("postgres", cfg.DBConnection)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	cpID := seedPocket(t, db, &Model{Name: "To be changed", Budget: 100.9, Balance: 0, Current: "THB", Description: "test pocket description", AccountId: 1})

	hPocket := New(db)
	e.DELETE("/cloud-pockets/:id", hPocket.Delete)
	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/cloud-pockets/%d", cpID), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	expected := `{"message": "Cloud pocket deleted successfully"}`
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, expected, rec.Body.String())
}
