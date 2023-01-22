//go:build integration

package cloudpocket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestUpdatePocketIT(t *testing.T) {
	db, e := setup(t)
	defer db.Close()
	id := seedPocket(t, db, &Model{Name: "To be changed", Budget: 100.9, Balance: 999.99, Current: "THB", Description: "test pocket description", AccountId: 1})
	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/cloud-pockets/%d", id), strings.NewReader(`{"name": "test_pocket"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	res := &Model{}
	err := json.Unmarshal(rec.Body.Bytes(), res)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, res.ID)
	assert.Equal(t, "test_pocket", res.Name)
	assert.Equal(t, http.StatusOK, rec.Code)
}
