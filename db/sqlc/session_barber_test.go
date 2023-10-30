package db

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/require"
// )

// func TestCreateSessionBabrer(t *testing.T) {
// 	arg := CreateSessionBarberParams{
// 		ID:           uuid.New(),
// 		Username:     "ThinhVan",
// 		RefreshToken: "123",
// 		UserAgent:    "foo",
// 		ClientIp:     "127.0.0.1",
// 		FcmDevice:    "127.0.0.1",
// 		IsBlocked:    false,
// 		ExpiresAt:    time.Now().Add(10 * time.Minute),
// 	}
// 	sessionBarber, err := testQueries.CreateSessionBarber(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, sessionBarber)
// }

// func TestGetSession(t *testing.T) {
// 	idSession := "10af67b7-0125-4870-ab8f-33cc2fa8d049"
// 	sessionBarber, err := testQueries.GetSession(context.Background(), uuid.MustParse(idSession))
// 	require.NoError(t, err)
// 	require.NotEmpty(t, sessionBarber)
// }
