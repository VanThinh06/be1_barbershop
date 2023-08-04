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

func TestGetListStore(t *testing.T) {
	arg := GetListStoreParams{
		Limit:  5,
		Offset: 0,
	}
	listStore, err := testQueries.GetListStore(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, listStore)
}
