package db

import (
	"context"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestCreateBarber(t *testing.T) {
	arg := CreateBarberParams{
		FullName:       "DuongVanThinh",
		Email:          "dvanthinh@gmail.com",
		HashedPassword: "VanThinh123",
		Role:           1,
		NickName: "DuongVanThinh",
		Phone: "0393482512",
		Gender: 1,

	}
	barber, err := testQueries.CreateBarber(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, barber)
}

func TestGetBarber(t *testing.T) {
	barber, err := testQueries.GetEmailBarber(context.Background(), "dvanthinh@gmail.com")
	require.NoError(t, err)
	require.NotEmpty(t, barber)
}
