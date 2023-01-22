package cloudpocket

import (
	"database/sql"
	"testing"

	"github.com/kkgo-software-engineering/workshop/config"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (*sql.DB, *echo.Echo) {
	cfg := config.New().All()
	db, err := sql.Open("postgres", cfg.DBConnection)
	assert.NoError(t, err)

	hPocket := New(db)
	e := echo.New()
	e.PUT("/cloud-pockets/:id", hPocket.HandleUpdatePocket)
	e.POST("/cloud-pockets", hPocket.HandleCreatePocket)

	return db, e
}

func seedPocket(t *testing.T, db *sql.DB, c *Model) int64 {
	err := db.QueryRow(`INSERT INTO cloud_pockets
	(name, budget, balance, is_default, description, currency, account_id) 
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	RETURNING id`,
		c.Name, c.Budget, c.Balance, c.IsDefault, c.Description, c.Current, c.AccountId).Scan(&c.ID)
	assert.NoError(t, err)
	return c.ID
}
