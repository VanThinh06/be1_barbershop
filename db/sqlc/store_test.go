package db

import (
	"barbershop/db/util"
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateStore(t *testing.T) {
	arg := CreateStoreParams{
		NameID:    util.RandomString(5),
		NameStore: util.RandomName(),
		Address:   util.RandomName(),
		Location:  0.999,
		Status:    1,
		ManagerID: []string{"VanThinhDuong"},
	}

	store, err := testQueries.CreateStore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, store)
}
