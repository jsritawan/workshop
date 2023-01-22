package cloudpocket

import (
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreatePocket(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	assert.NoError(t, err)

	mock.ExpectQuery("INSERT INTO cloud_pockets.*").
		WithArgs("cloud pocket name", 0.0, 100.0, false, "", "THB", 1).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	handler := New(db)

	err = handler.CreatePocket(&Model{
		Name:      "cloud pocket name",
		Balance:   100.0,
		Current:   "THB",
		AccountId: 1,
	})
	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestCreatePocketError(t *testing.T) {
	db, mock, err := sqlmock.New()
	defer db.Close()
	assert.NoError(t, err)

	mock.ExpectQuery("INSERT INTO cloud_pockets.*").
		WithArgs("cloud pocket name", 0.0, 100.0, false, "", "THB", 1).
		WillReturnError(echo.NewHTTPError(http.StatusInternalServerError, "failed to create cloud pocket"))

	handler := New(db)

	err = handler.CreatePocket(&Model{
		Name:      "cloud pocket name",
		Balance:   100.0,
		Current:   "THB",
		AccountId: 1,
	})
	assert.Error(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
