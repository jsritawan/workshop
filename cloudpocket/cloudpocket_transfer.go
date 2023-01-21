package cloudpocket

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Resp struct {
	PocketID int     `json:"pocketId"`
	Balance  float64 `json:"balance"`
}

type Res struct {
	PocketID int     `json:"pocketId"`
	Amount   float64 `json:"amount"`
}

func (h handler) Transfer(echo echo.Context) error {

	req := Res{}
	id, err := strconv.Atoi(echo.Param("id"))

	if err != nil {
		return echo.JSON(http.StatusBadRequest, "Invalid param id")
	}

	err = echo.Bind(&req)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, "Invalid request body")
	}

	var fromBalance float64
	row := h.db.QueryRow(cStmtGetBalance, id)
	if err := row.Scan(&fromBalance); err != nil {
		return echo.JSON(http.StatusInternalServerError, "Not found my pocket from id")
	}

	if fromBalance < req.Amount {
		return echo.JSON(http.StatusInternalServerError, "Not enough balance")
	}

	stmt, err := h.db.Prepare(cStmtUpdateBalance)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, "prepare sql error from id")
	}

	if _, err := stmt.Exec(fromBalance-req.Amount, id); err != nil {
		return echo.JSON(http.StatusInternalServerError, "update balance error 1")
	}

	var toBalance float64
	row = h.db.QueryRow(cStmtGetBalance, req.PocketID)
	if err := row.Scan(&toBalance); err != nil {
		return echo.JSON(http.StatusInternalServerError, "prepare sql error pocket id")
	}

	if _, err := stmt.Exec(toBalance+req.Amount, req.PocketID); err != nil {
		return echo.JSON(http.StatusInternalServerError, "update balance error 2")
	}

	return echo.JSON(200, Resp{
		PocketID: req.PocketID,
		Balance:  toBalance + req.Amount,
	})
}
