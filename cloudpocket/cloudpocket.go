package cloudpocket

import (
	"database/sql"
)

type handler struct {
	db *sql.DB
}

type Model struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Budget      float64 `json:"budget"`
	Balance     float64 `json:"balance"`
	IsDefault   bool    `json:"isDefault"`
	Description string  `json:"description"`
	Current     string  `json:"current"`
	AccountId   int64   `json:"accountId"`
}

func New(db *sql.DB) *handler {
	return &handler{db}
}

const (
	cStmtUpdateBalance = "UPDATE cloud_pockets SET balance = $1 WHERE id = $2;"
	cStmtGetBalance    = "SELECT balance FROM cloud_pockets WHERE id = $1;"
)
