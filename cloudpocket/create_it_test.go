//go:build intergration

package cloudpocket

import (
	"database/sql"
	"testing"

	"github.com/kkgo-software-engineering/workshop/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreatePocketIT(t *testing.T) {
	e := echo.New()

	cfg := config.New().All()
	sql, err := sql.Open("postgres", cfg.DBConnection)
	assert.NoError(t, err)

	hPocket := New(sql)

	e.POST("/cloud-pockets", hPocket.HandleCreatePocket)
	// req := httptest.NewRequest(http.MethodPost, "/cloud-pockets", strings.NewReader(`{"name": "test_pocket", "balance": 999.99, "currency": "THB", "description": "test pocket description", "account_id": 1}`))
}
