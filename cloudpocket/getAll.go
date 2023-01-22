package cloudpocket

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Err struct {
	Message string `json:"message"`
}

func (h *handler) GetAll(c echo.Context) error {
	cloudpockets := []Model{}
	// statement, err := h.db.Prepare("SELECT id, name, budget, balance, is_default, description, currency, account_id FROM cloud_pockets")
	statement, err := h.db.Prepare("SELECT * FROM cloud_pockets")
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query all"})
	}

	rows, err := statement.Query()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't query all"})
	}

	for rows.Next() {
		var cp Model
		err = rows.Scan(&cp.ID, &cp.Name, &cp.Budget, &cp.Balance, &cp.IsDefault, &cp.Description, &cp.Current, &cp.AccountId)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan pockets - " + err.Error()})
		}
		cloudpockets = append(cloudpockets, cp)
	}
	return c.JSON(http.StatusOK, cloudpockets)

}
