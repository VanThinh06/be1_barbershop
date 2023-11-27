package token

import (
	"time"

	"github.com/google/uuid"
)

type Maker interface {
	CreateToken(Barber BarberPayload, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
	VerifyTokenCustomer(token string) (*CustomerPayload, error)
}

type CustomerMaker interface {
	CreateToken(customer Customer, duration time.Duration) (string, *CustomerPayload, error)
	VerifyToken(token string) (*CustomerPayload, error)
	RefreshToken(id uuid.UUID, customer Customer, duration time.Duration) (string, *CustomerPayload, error)
}
