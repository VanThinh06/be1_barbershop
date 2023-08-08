// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateBarber(ctx context.Context, arg CreateBarberParams) (Barber, error)
	CreateService(ctx context.Context, arg CreateServiceParams) (Service, error)
	CreateServiceCategory(ctx context.Context, arg CreateServiceCategoryParams) (ServiceCategory, error)
	CreateSessionBarber(ctx context.Context, arg CreateSessionBarberParams) (SessionsBarber, error)
	CreateStore(ctx context.Context, arg CreateStoreParams) (Store, error)
	DeleteServiceCategory(ctx context.Context, id uuid.UUID) (ServiceCategory, error)
	DeleteServicewithStoreCategory(ctx context.Context, serviceCategoryID uuid.UUID) error
	GetBarber(ctx context.Context, username string) (Barber, error)
	GetListServiceCategorywithStore(ctx context.Context, storeID uuid.UUID) ([]ServiceCategory, error)
	GetListServicewithCategory(ctx context.Context, arg GetListServicewithCategoryParams) ([]Service, error)
	GetListStore(ctx context.Context, arg GetListStoreParams) ([]Store, error)
	GetService(ctx context.Context, id uuid.UUID) (Service, error)
	GetServiceCategory(ctx context.Context, id uuid.UUID) (ServiceCategory, error)
	GetSession(ctx context.Context, id uuid.UUID) (SessionsBarber, error)
	GetStore(ctx context.Context, id uuid.UUID) (Store, error)
}

var _ Querier = (*Queries)(nil)
