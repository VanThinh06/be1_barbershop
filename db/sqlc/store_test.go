package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"
)

func TestCreateRandomEmployee(t *testing.T) {
	arg := CreateStoreParams{
		NameStore:  null.String{},
		ManagerID:  uuid.MustParse("4714dc3f-2f14-4abe-b19f-d8b925334b48"),
		EmployeeID: uuid.MustParse("1792c33c-4102-4756-b27d-ab519e192275"),
		Location:   sql.NullInt32{Int32: 1, Valid: false},
		Status:     null.String{},
	}

	employee, err := testQueries.CreateStore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, employee)

	require.Equal(t, employee.NameStore, arg.NameStore)
	require.Equal(t, employee.Status, arg.Status)

}
