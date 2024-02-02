package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Barber struct {
	BarberID       uuid.UUID `json:"barber_id"`
	NickName       string    `json:"nick_name"`
	BarberRole     int32     `json:"role"`
	BarberRoleType string    `json:"role_type"`
	Phone          string    `json:"phone"`
	Email          string    `json:"email"`
	FcmDevice      string    `json:"fcm_device"`
}

type BarberPayload struct {
	ID        uuid.UUID `json:"id"`
	Barber    Barber    `json:"barber"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expired_at"`
}

func NewPayload(barber Barber, duration time.Duration) (*BarberPayload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &BarberPayload{
		ID:        tokenID,
		Barber:    barber,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}
	return payload, nil
}

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

func RePayloadBarber(id uuid.UUID, barber Barber, duration time.Duration) (*BarberPayload, error) {
	payload := &BarberPayload{
		ID:        id,
		Barber:    barber,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *BarberPayload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ErrExpiredToken
	}
	return nil
}
