package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateBarber(t *testing.T) {
	users := CreateRandomUsers(t)

	arg := CreateBarberParams{
		NameID:       users.Username,
		StoreID:      uuid.NullUUID{},
		StoreManager: nil,
	}

	barber, err := testQueries.CreateBarber(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, barber)
	require.Equal(t, arg.NameID, barber.NameID)
	require.Equal(t, arg.StoreID, barber.StoreID)
	require.Equal(t, arg.StoreManager, barber.StoreManager)
	require.NotZero(t, barber.CreatedAt)
}
