package token

import (
	// "barbershop/src/shared/utils"
	// "testing"
	// "time"

	// "github.com/google/uuid"
	// "github.com/stretchr/testify/require"
)

// func TestPasetoMaker(t *testing.T) {
// 	maker, err := NewPasetoMaker(utils.RandomString(32))
// 	require.NoError(t, err)

// 	barber_id, err := uuid.NewUUID()
// 	if err != nil {
// 		return
// 	}
// 	duration := time.Minute

// 	issuedAt := time.Now()
// 	expiredAt := issuedAt.Add(duration)

// 	token, payload, err := maker.CreateToken(, duration)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, token)
// 	require.NotEmpty(t, payload)

// 	payload, err = maker.VerifyToken(token)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, token)

// 	require.NotZero(t, payload.ID)
// 	require.Equal(t, barber_id, payload.Barber)
// 	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
// 	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
// }

// func TestExpiredPasetoToken(t *testing.T) {
// 	maker, err := NewPasetoMaker(utils.RandomString(32))
// 	require.NoError(t, err)
// 	barber_id, err := uuid.NewUUID()
// 	if err != nil {
// 		return
// 	}

// 	token, payload, err := maker.CreateToken(barber_id, -time.Minute)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, token)
// 	require.NotEmpty(t, payload)

// 	payload, err = maker.VerifyToken(token)
// 	require.Error(t, err)
// 	require.EqualError(t, err, ErrExpiredToken.Error())
// 	require.Nil(t, payload)
// }
