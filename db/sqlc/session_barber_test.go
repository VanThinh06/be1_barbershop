package db

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateSessionBabrer(t *testing.T) {
	arg := CreateSessionBarberParams{
		Username:     "ThinhVan",
		RefreshToken: "123",
		UserAgent:    "foo",
		IsManager:    false,
		ClientIp:     "127.0.0.1",
		FcmDevice:    "127.0.0.1",
		IsBlocked:    false,
		ExpiresAt:    time.Now().Add(10 * time.Minute),
	}
	sessionBarber, err := testQueries.CreateSessionBarber(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, sessionBarber)
}

func TestGetSession(t *testing.T) {
	idSession := "d78e9aae-8ed2-4d51-80c8-ce0644dd813c"
	sessionBarber, err := testQueries.GetSession(context.Background(), uuid.MustParse(idSession))
	require.NoError(t, err)
	require.NotEmpty(t, sessionBarber)
}
