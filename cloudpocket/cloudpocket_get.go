package cloudpocket

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetCloudpocketByID(c echo.Context) error {
	id := c.Param("id")

	stmt, err := h.db.Prepare("SELECT id, name, budget, balance, is_default, description, currency, account_id FROM cloud_pockets WHERE id=$1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query statment with id: " + err.Error()})
	}

	row := stmt.QueryRow(id)
	cp := Model{}
	err = row.Scan(&cp.ID, &cp.Name, &cp.Budget, &cp.Balance, &cp.IsDefault, &cp.Description, &cp.Current, &cp.AccountId)

	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "cloud_pockets not found: " + err.Error()})
	case nil:
		return c.JSON(http.StatusOK, cp)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan: " + err.Error()})
	}
}

func (h *handler) GetCloudpocket(c echo.Context) error {
	stmt, err := h.db.Prepare("SELECT id, name, budget, balance, is_default, description, currency, account_id FROM cloud_pockets")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query all statment: " + err.Error()})
	}

	rows, err := stmt.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't query all: " + err.Error()})
	}

	cps := []Model{}
	for rows.Next() {
		cp := Model{}
		err := rows.Scan(&cp.ID, &cp.Name, &cp.Budget, &cp.Balance, &cp.IsDefault, &cp.Description, &cp.Current, &cp.AccountId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan: " + err.Error()})
		}

		cps = append(cps, cp)
	}

	return c.JSON(http.StatusOK, cps)
}
