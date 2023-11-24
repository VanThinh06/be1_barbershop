package token

import (
	"time"

)

type Maker interface {
	CreateToken(Barber BarberPayload, duration time.Duration) (string, *Payload, error)
	VerifyToken(token string) (*Payload, error)
}
