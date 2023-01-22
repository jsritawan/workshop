package cloudpocket

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/kkgo-software-engineering/workshop/mlog"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type DeleteResponse struct {
	Message string `json:"message"`
}

func (h handler) Delete(c echo.Context) error {
	logger := mlog.L(c)
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("bad request body", zap.Error(err))
		return c.JSON(http.StatusBadRequest, DeleteResponse{"Cannot convert cloud pocket id to int"})
	}

	bRow := h.db.QueryRowContext(ctx, cStmtGetBalance, id)
	if bRow.Err() == sql.ErrNoRows {
		logger.Error("Cloud pocket not found", zap.Error(bRow.Err()))
		return c.JSON(http.StatusNotFound, DeleteResponse{"Cloud pocket not found"})
	}

	var balance float64
	if err := bRow.Scan(&balance); err != nil {
		logger.Error("Cannot scan balance", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, DeleteResponse{"Cannot scan cloud pocket balance"})
	}

	if balance > 0 {
		logger.Error("bad request body")
		return c.JSON(http.StatusBadRequest, DeleteResponse{"Unable to delete this Cloud Pocket\n there is amount left in this Cloud Pocket, please move money out and try again"})
	}

	dRow := h.db.QueryRowContext(ctx, cStmtDelete, id)
	var dID int

	if err := dRow.Scan(&dID); err != nil {
		logger.Error("Cannot scan id", zap.Error(err))
		return c.JSON(http.StatusNotFound, DeleteResponse{"Cannot scan cloud pocket ID"})
	}

	if dID == 0 {
		logger.Error("Cannot delete cloud pocket")
		c.JSON(http.StatusInternalServerError, DeleteResponse{"Cannot delete this Cloud pocket"})
	}

	logger.Info("create successfully", zap.Int("id", dID))
	return c.JSON(http.StatusOK, DeleteResponse{"Cloud pocket deleted successfully"})
}
