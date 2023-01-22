package cloudpocket

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
)

type Resp struct {
	PocketID int     `json:"pocketId"`
	Balance  float64 `json:"balance"`
}

type Res struct {
	PocketID int     `json:"pocketId"`
	Amount   float64 `json:"amount"`
}

func addFund(current, newAmount float64) float64 {
	var a = decimal.NewFromFloat(current).Add(decimal.NewFromFloat(newAmount))
	num, err := a.Float64()
	if err != true {
		fmt.Println(err)
	}
	return num
}

func deleteFund(current, newAmount float64) float64 {
	var a = decimal.NewFromFloat(current).Sub(decimal.NewFromFloat(newAmount))
	num, err := a.Float64()
	if err != true {
		fmt.Println(err)
	}
	return num
}

func (h handler) Transfer(echo echo.Context) error {
	ctx := echo.Request().Context()
	req := Res{}
	id, err := strconv.Atoi(echo.Param("id"))

	if err != nil {
		return echo.JSON(http.StatusBadRequest, "Invalid param id")
	}

	err = echo.Bind(&req)

	if err != nil {
		return echo.JSON(http.StatusBadRequest, "Invalid request body")
	}

	tx, err := h.db.BeginTx(ctx, nil)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, err.Error())
	}

	var fromBalance float64
	row := tx.QueryRowContext(ctx, cStmtGetBalance, id)
	if err := row.Scan(&fromBalance); err != nil {
		return echo.JSON(http.StatusInternalServerError, "Not found my pocket from id")
	}

	if fromBalance < req.Amount {
		return echo.JSON(http.StatusInternalServerError, "Not enough balance")
	}

	stmt, err := h.db.PrepareContext(ctx, cStmtUpdateBalance)

	if err != nil {
		return echo.JSON(http.StatusInternalServerError, "prepare sql error from id")
	}

	if _, err := stmt.ExecContext(ctx, deleteFund(fromBalance, req.Amount), id); err != nil {
		tx.Rollback()
	}

	var toBalance float64
	row = h.db.QueryRowContext(ctx, cStmtGetBalance, req.PocketID)
	if err := row.Scan(&toBalance); err != nil {
		return echo.JSON(http.StatusInternalServerError, "prepare sql error pocket id")
	}

	if _, err := stmt.ExecContext(ctx, addFund(toBalance, req.Amount), req.PocketID); err != nil {
		tx.Rollback()
	}
	if err = tx.Commit(); err != nil {
		return echo.JSON(http.StatusInternalServerError, Err{Message: "transaction failed"})
	}

	return echo.JSON(200, Resp{
		PocketID: req.PocketID,
		Balance:  addFund(toBalance, req.Amount),
	})
}
