package cloudpocket

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) HandleCreatePocket(c echo.Context) error {
	pocket := &Model{}
	if err := c.Bind(pocket); err != nil || pocket.AccountId == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	err := h.CreatePocket(pocket)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, pocket)
}

func (h *handler) CreatePocket(c *Model) error {
	err := h.db.QueryRow(`INSERT INTO cloud_pockets
	(name, budget, balance, is_default, description, currency, account_id) 
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	RETURNING id`,
		c.Name, c.Budget, c.Balance, c.IsDefault, c.Description, c.Current, c.AccountId).Scan(&c.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create cloud pocket"+err.Error())
	}
	return nil
}
