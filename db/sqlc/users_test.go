package db

import (
	"barbershop/db/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// tên file test phải kết thúc bằng hậu tố _test

func TestListEmployee(t *testing.T) {
	for i := 0; i < 5; i++ {
		CreateRandomUsers(t)
	}
}

func CreateRandomUsers(t *testing.T) User {
	arg := CreateUsersParams{
		Username:       util.RandomName(),
		FullName:       util.RandomString(15),
		Email:          util.RandomEmail(),
		HashedPassword: "secret",
		FcmDevice:      "oke",
	}
	users, err := testQueries.CreateUsers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)

	require.Equal(t, users.Username, arg.Username)
	require.Equal(t, users.FullName, arg.FullName)
	require.Equal(t, users.Email, arg.Email)
	require.NotZero(t, users.CreatedAt)
	require.True(t, users.PasswordChangedAt.IsZero())

	// delete user
	testQueries.DeleteUsers(context.Background(), arg.Username)
	return users
}

func TestGetListEmployee(t *testing.T) {
	users := CreateRandomUsers(t)
	arg, err := testQueries.GetUsers(context.Background(), users.Username)
	require.NoError(t, err)
	require.NotEmpty(t, arg)
	require.Equal(t, users.Username, arg.Username)
	require.Equal(t, users.FullName, arg.FullName)
	require.Equal(t, users.Email, arg.Email)
	require.NotZero(t, users.CreatedAt)
	require.WithinDuration(t, users.PasswordChangedAt, arg.PasswordChangedAt, time.Second)
	require.WithinDuration(t, users.CreatedAt, arg.CreatedAt, time.Second)
}
