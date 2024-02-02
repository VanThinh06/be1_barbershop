package token

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	CustomerID uuid.UUID      `json:"customer_id"`
	Name       string         `json:"name"`
	Phone      sql.NullString `json:"phone"`
	Email      string         `json:"email"`
	FcmDevice  string         `json:"fcm_device"`
}

type CustomerPayload struct {
	ID        uuid.UUID `json:"id"`
	Customer  Customer  `json:"customer"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expired_at"`
}

func NewPayloadCustomer(customer Customer, duration time.Duration) (*CustomerPayload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &CustomerPayload{
		ID:        tokenID,
		Customer:  customer,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}
	return payload, nil
}

func RePayloadCustomer(id uuid.UUID, customer Customer, duration time.Duration) (*CustomerPayload, error) {
	payload := &CustomerPayload{
		ID:        id,
		Customer:  customer,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *CustomerPayload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return ErrExpiredToken
	}
	return nil
}
