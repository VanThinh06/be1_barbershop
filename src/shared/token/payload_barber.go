package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)
type BarberPayload struct {
	BarberID uuid.UUID `json:"username"`
	Role     int32     `json:"role"`
	Phone    string    `json:"phone"`
	Email    string    `json:"email"`
}

type Payload struct {
	ID        uuid.UUID     `json:"id"`
	Barber    BarberPayload `json:"barber"`
	IssuedAt  time.Time     `json:"issued_at"`
	ExpiredAt time.Time     `json:"expired_at"`
}

func NewPayload(barber BarberPayload, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Barber:  barber,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
