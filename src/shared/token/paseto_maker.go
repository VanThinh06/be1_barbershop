package token

import (
	"barbershop/src/shared/utils"
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username BarberPayload, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, payload, err
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
func (maker *PasetoMaker) VerifyTokenCustomer(token string) (*CustomerPayload, error) {
	payload := &CustomerPayload{}
	config, err := utils.LoadConfig(".")
	if err != nil {
		return nil, ErrInvalidToken
	}
	
	maker = &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(config.TokenSymmetricKeyCustomer),
	}

	err = maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// / paseto customer
type PasetoMakerCustomer struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMakerCustomer(symmetricKey string) (CustomerMaker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMakerCustomer{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (maker *PasetoMakerCustomer) CreateToken(customer Customer, duration time.Duration) (string, *CustomerPayload, error) {
	payload, err := NewPayloadCustomer(customer, duration)
	if err != nil {
		return "", payload, err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, payload, err
}

func (maker *PasetoMakerCustomer) VerifyToken(token string) (*CustomerPayload, error) {
	payload := &CustomerPayload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (maker *PasetoMakerCustomer) RefreshToken(id uuid.UUID, customer Customer, duration time.Duration) (string, *CustomerPayload, error) {
	payload, err := RePayloadCustomer(id, customer, duration)
	if err != nil {
		return "", payload, err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, payload, err
}
