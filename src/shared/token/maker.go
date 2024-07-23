package token

import (
	"time"

	"github.com/google/uuid"
)

type Maker interface {
	CreateToken(Barber Barber, duration time.Duration) (string, *BarberPayload, error)
	VerifyToken(token string) (*BarberPayload, error)
	RefreshToken(id uuid.UUID, barber Barber, duration time.Duration) (string, *BarberPayload, error)
	VerifyTokenCustomer(token string) (*CustomerPayload, error)
	CreateAESString(plaintext string) (string, error)
	DecodeAESString(encryptedText string) (string, error)
}

type CustomerMaker interface {
	CreateToken(customer Customer, duration time.Duration) (string, *CustomerPayload, error)
	VerifyToken(token string) (*CustomerPayload, error)
	RefreshToken(id uuid.UUID, customer Customer, duration time.Duration) (string, *CustomerPayload, error)
}
