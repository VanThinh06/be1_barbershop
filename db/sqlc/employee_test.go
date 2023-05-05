package db

import (
	"barbershop/db/util"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"
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
		Image:    null.String{},
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
	require.NoError(t, err)
	require.NotEmpty(t, listEmployee)

}

func TestUpdateEmployee(t *testing.T) {
	employee, err := testQueries.GetEmployee(context.Background(), uuid.MustParse("00489b51-3966-430d-9cb0-294eebbeca7d"))
	require.NoError(t, err)
	require.NotEmpty(t, employee)

	arg := UpdateEmployeeParams{
		ID:       uuid.MustParse("00489b51-3966-430d-9cb0-294eebbeca7d"),
		Username: "Okela",
		Image:    employee.Image,
		StoreID:  uuid.MustParse("00000000-0000-0000-0000-000000000000"),
	}

	employeeUpdate, err := testQueries.UpdateEmployee(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, employeeUpdate)

}

func TestDeleteEmployee(t *testing.T) {
	err := testQueries.DeleteEmployee(context.Background(), uuid.MustParse("18a072c7-f86d-4bb7-bd36-36580ec07fec"))
	require.NoError(t, err)

}
