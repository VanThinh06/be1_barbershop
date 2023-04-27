package db

import (
	"barbershop/db/util"
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

// tên file test phải kết thúc bằng hậu tố _test

func TestListEmployee(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEmployee(t)
	}
}

func createRandomEmployee(t *testing.T) {
	arg := CreateEmployeeParams{
		Username: util.RandomName(),
		Role:     util.RandomRole(),
		Image:    sql.NullString{Valid: true},
		StoreID:  uuid.Nil,
	}

	employee, err := testQueries.CreateEmployee(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, employee)

	require.Equal(t, employee.Username, arg.Username)
	require.Equal(t, employee.Role, arg.Role)
	require.Equal(t, employee.Image, arg.Image)
	require.NotZero(t, employee.CreatedAt)

}

func TestGetListEmployee(t *testing.T) {

	arg := GetListEmployeewithStoreParams{
		StoreID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
		Limit:   10,
	}

	listEmployee, err := testQueries.GetListEmployeewithStore(context.Background(), arg)

	fmt.Println(listEmployee)
	require.NoError(t, err)
	require.NotEmpty(t, listEmployee)

}

func TestUpdateEmployee(t *testing.T) {
	employee, err := testQueries.GetEmployee(context.Background(), uuid.MustParse("18a072c7-f86d-4bb7-bd36-36580ec07fec"))
	require.NoError(t, err)
	require.NotEmpty(t, employee)

	arg := UpdateAuthorParams{
		ID:       uuid.MustParse("18a072c7-f86d-4bb7-bd36-36580ec07fec"),
		Username: "Okela",
		Image:    employee.Image,
		StoreID:  employee.StoreID,
	}

	employeeUpdate, err := testQueries.UpdateAuthor(context.Background(), arg)

	fmt.Println(employeeUpdate)
	require.NoError(t, err)
	require.NotEmpty(t, employeeUpdate)

}

func TestDeleteEmployee(t *testing.T) {
	err := testQueries.DeleteEmployee(context.Background(), uuid.MustParse("18a072c7-f86d-4bb7-bd36-36580ec07fec"))
	require.NoError(t, err)

}
