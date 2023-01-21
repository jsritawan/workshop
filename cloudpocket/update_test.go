package cloudpocket

import (
	"database/sql/driver"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUpdatePocket(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("UPDATE cloud_pockets.*").ExpectExec().
		WithArgs(1, "name", 100.0, "description").
		WillReturnResult(driver.RowsAffected(1))

	handler := New(db)

	err = handler.UpdatePocket(&Model{
		ID:          1,
		Name:        "name",
		Budget:      100.0,
		Description: "description",
	})
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUpdatePocketNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("UPDATE cloudpocket.*").ExpectExec().
		WithArgs(1, "name", 100.0, "description").
		WillReturnResult(driver.RowsAffected(0))

	handler := New(db)

	err = handler.UpdatePocket(&Model{
		ID:          1,
		Name:        "name",
		Budget:      100.0,
		Description: "description",
	})
	assert.Error(t, err)
}
