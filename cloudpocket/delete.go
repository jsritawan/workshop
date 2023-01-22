package cloudpocket

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h handler) Delete(c echo.Context) error {
	return c.JSON(http.StatusInternalServerError, `{"message": "not implemented yet"}`)
}
