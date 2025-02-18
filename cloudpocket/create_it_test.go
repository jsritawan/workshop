//go:build integration

package cloudpocket

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreatePocketIT(t *testing.T) {
	db, e := setup(t)
	defer db.Close()
	req := httptest.NewRequest(http.MethodPost, "/cloud-pockets", strings.NewReader(`{"name": "test_pocket", "balance": 999.99, "current": "THB", "description": "test pocket description", "accountId": 1}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	res := &Model{}
	err := json.Unmarshal(rec.Body.Bytes(), res)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, res.ID)
	assert.Equal(t, "test_pocket", res.Name)
	assert.Equal(t, 999.99, res.Balance)
	assert.Equal(t, "test pocket description", res.Description)
	assert.Equal(t, "THB", res.Current)
	assert.Equal(t, int64(1), res.AccountId)
	assert.Equal(t, http.StatusCreated, rec.Code)
}
