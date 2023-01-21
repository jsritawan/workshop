package cloudpocket

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *handler) HandleUpdatePocket(c echo.Context) error {
	var pocket CloudPocket
	err := c.Bind(&pocket)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	param := c.Param("id")
	pocket.ID, err = strconv.ParseInt(param, 10, 64)
	if pocket.ID == 0 || err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	if err := h.UpdatePocket(&pocket); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, pocket)
}

func (h *handler) UpdatePocket(c *CloudPocket) error {
	stmt, err := h.db.Prepare(
		`UPDATE cloudpocket
		SET
			name=$2,
			budget=$3,
			description=$4
		WHERE id=$1 `,
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot prepare statement")
	}

	var res sql.Result
	res, err = stmt.Exec(c.ID, c.Name, c.Budget, c.Description)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot update pocketss: "+err.Error())
	}

	row, err := res.RowsAffected()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "cannot update pocket: "+err.Error())
	}
	if row == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "pocket not found")
	}
	return nil
}
