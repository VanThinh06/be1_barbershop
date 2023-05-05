package db

import (
	"barbershop/db/util"
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"gopkg.in/guregu/null.v4"
)

func TestCreateManager(t *testing.T) {
	arg := CreateManagerParams{
		Username: util.RandomName(),
		Role:     util.RandomRole(),
		Image:    null.String{},
	}

	employee, err := testQueries.CreateManager(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, employee)
}

func TestUpdateManager(t *testing.T) {
	arg := UpdateManagerParams{
		Username: util.RandomName(),
		Role:     util.RandomRole(),
		Image:    null.String{},
		ID:       uuid.MustParse("4714dc3f-2f14-4abe-b19f-d8b925334b48"),
	}
	manager, err := testQueries.UpdateManager(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, manager)
}
