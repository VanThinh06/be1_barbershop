// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateBarber(ctx context.Context, arg CreateBarberParams) (Barber, error)
	CreateUsers(ctx context.Context, arg CreateUsersParams) (User, error)
	GetBarber(ctx context.Context, nameID string) (Barber, error)
	GetUsers(ctx context.Context, username string) (User, error)
}

var _ Querier = (*Queries)(nil)
