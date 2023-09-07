package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"
)

func TestCreateBarber(t *testing.T) {
	arg := CreateBarberParams{
		Username:       "VanThinhDuong",
		FullName:       "DuongVanThinh",
		Email:          "dvanthinh@gmail.com",
		HashedPassword: "VanThinh123",
		Avatar:         null.StringFrom("https://"),
		Role:           null.StringFrom(""),
		Status:         null.StringFrom(""),
		StoreWork:      uuid.NullUUID{},
	}
	barber, err := testQueries.CreateBarber(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, barber)
}

func TestGetBarber(t *testing.T) {
	nameBarber := "ThinhVan"
	barber, err := testQueries.GetBarber(context.Background(), nameBarber)
	require.NoError(t, err)
	require.NotEmpty(t, barber)
}
