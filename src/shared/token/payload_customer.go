package token

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	CustomerID uuid.UUID `json:"username"`
	Gender     int32     `json:"gender"`
	Phone      sql.NullString    `json:"phone"`
	Email      string    `json:"email"`
	FcmDevice  string    `json:"fcm_device"`
	Timezone   string    `json:"timezone"`
}

type CustomerPayload struct {
	ID        uuid.UUID `json:"id"`
	Customer  Customer  `json:"customer"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
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
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func RePayloadCustomer(id uuid.UUID, customer Customer, duration time.Duration) (*CustomerPayload, error) {
	payload := &CustomerPayload{
		ID:        id,
		Customer:  customer,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func (payload *CustomerPayload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
