package token

import (
	// "barbershop/src/shared/utils"
	// "testing"
	// "time"

	// "github.com/dgrijalva/jwt-go"
	// "github.com/google/uuid"
	// "github.com/stretchr/testify/require"
)

// func TestJWTMaker(t *testing.T) {
// 	maker, err := NewJWTMaker(utils.RandomString(32))
// 	require.NoError(t, err)

// 	barber_id, err := uuid.NewUUID()
// 	if err != nil {
// 		return
// 	}
// 	duration := time.Minute

// 	issuedAt := time.Now()
// 	expiredAt := issuedAt.Add(duration)

// 	token, payload, err := maker.CreateToken(barber_id, duration)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, token)

// 	payload, err = maker.VerifyToken(token)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, token)

// 	require.NotZero(t, payload.ID)

// 	require.Equal(t, barber_id, payload.BarberID)
// 	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
// 	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
// }

// func TestExpiredJWTToken(t *testing.T) {
// 	maker, err := NewJWTMaker(utils.RandomString(32))
// 	require.NoError(t, err)

// 	barber_id, err := uuid.NewUUID()
// 	if err != nil {
// 		return
// 	}

// 	token, payload, err := maker.CreateToken(barber_id, -time.Minute)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, token)

// 	payload, err = maker.VerifyToken(token)
// 	require.Error(t, err)
// 	require.EqualError(t, err, ErrExpiredToken.Error())
// 	require.Nil(t, payload)
// }

// func TestInvalidJWTTokenAlgNone(t *testing.T) {
// 	barber_id, err := uuid.NewUUID()
// 	if err != nil {
// 		return
// 	}
// 	payload, err := NewPayload(barber_id, time.Minute)
// 	require.NoError(t, err)

// 	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
// 	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
// 	require.NoError(t, err)

// 	maker, err := NewJWTMaker(utils.RandomString(32))
// 	require.NoError(t, err)

// 	payload, err = maker.VerifyToken(token)
// 	require.Error(t, err)
// 	require.EqualError(t, err, ErrInvalidToken.Error())
// 	require.Nil(t, payload)
// }
