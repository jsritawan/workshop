package cloudpocket

import (
	"database/sql"
)

type handler struct {
	db *sql.DB
}

func New(db *sql.DB) *handler {
	return &handler{db}
}

const (
	cStmtUpdateBalance = "UPDATE cloud_pockets SET balance = $1 WHERE id = $2;"
	cStmtGetBalance    = "SELECT balance FROM cloud_pockets WHERE id = $1;"

	cStmtUpdateBalanceTest = "UPDATE `cloud_pockets` SET (.+) WHERE (.+)"
	cStmtGetBalanceTest    = "SELECT `balance` FROM `cloud_pockets` WHERE (.+)"
)
