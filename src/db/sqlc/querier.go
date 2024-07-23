// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	ChangePasswordCustomer(ctx context.Context, arg ChangePasswordCustomerParams) (Customer, error)
	CheckBarberRolePermission(ctx context.Context, arg CheckBarberRolePermissionParams) (bool, error)
	CreateBarberShopChains(ctx context.Context, arg CreateBarberShopChainsParams) (BarberShopChain, error)
	CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error)
	CreateServiceCategory(ctx context.Context, arg CreateServiceCategoryParams) (ServiceCategory, error)
	CreateServiceItem(ctx context.Context, arg CreateServiceItemParams) (ServiceItem, error)
	CreateServicePackage(ctx context.Context, arg CreateServicePackageParams) (ServicePackage, error)
	CreateServicePackageItem(ctx context.Context, arg CreateServicePackageItemParams) (ServicePackageItem, error)
	CreateSessionBarber(ctx context.Context, arg CreateSessionBarberParams) (SessionsBarber, error)
	CreateSessionsCustomer(ctx context.Context, arg CreateSessionsCustomerParams) (SessionsCustomer, error)
	// DELETE
	DeleteBarberById(ctx context.Context, id uuid.UUID) error
	DeleteBarberRoleById(ctx context.Context, id uuid.UUID) error
	DeleteBarberShopById(ctx context.Context, id uuid.UUID) error
	DeleteBarberShopChains(ctx context.Context, id uuid.UUID) error
	DeleteServiceCategories(ctx context.Context, id int16) error
	DeleteServiceItem(ctx context.Context, id uuid.UUID) error
	DeleteServicePackage(ctx context.Context, id uuid.UUID) error
	GetBarberByCredential(ctx context.Context, arg GetBarberByCredentialParams) (GetBarberByCredentialRow, error)
	GetBarberById(ctx context.Context, id uuid.UUID) (Barber, error)
	GetBarberShopById(ctx context.Context, id uuid.UUID) (ViewBarberShop, error)
	GetBarberShopChains(ctx context.Context, id uuid.UUID) (BarberShopChain, error)
	GetBarberWithOptionalRole(ctx context.Context, arg GetBarberWithOptionalRoleParams) (GetBarberWithOptionalRoleRow, error)
	// QUERY SELECT
	GetBarberWithRoleAndShop(ctx context.Context, arg GetBarberWithRoleAndShopParams) (GetBarberWithRoleAndShopRow, error)
	GetCountOTPsForBarberToday(ctx context.Context, barberID uuid.UUID) (int64, error)
	GetCustomer(ctx context.Context, id uuid.UUID) (Customer, error)
	GetDistricts(ctx context.Context, provinceID int16) ([]District, error)
	GetOTPRequestDetails(ctx context.Context, arg GetOTPRequestDetailsParams) (GetOTPRequestDetailsRow, error)
	GetPermissionsByBarberShop(ctx context.Context, arg GetPermissionsByBarberShopParams) ([]Permission, error)
	GetProvinces(ctx context.Context) ([]Province, error)
	GetServiceItem(ctx context.Context, id uuid.UUID) (GetServiceItemRow, error)
	GetServicePackage(ctx context.Context, id uuid.UUID) (ViewServicePackage, error)
	GetSessionBarber(ctx context.Context, id uuid.UUID) (SessionsBarber, error)
	GetSessionsCustomer(ctx context.Context, id uuid.UUID) (SessionsCustomer, error)
	GetTimerServiceItem(ctx context.Context, dollar_1 []uuid.UUID) (int64, error)
	GetUserCustomer(ctx context.Context, arg GetUserCustomerParams) (Customer, error)
	GetWards(ctx context.Context, districtID int16) ([]Ward, error)
	InsertBarber(ctx context.Context, arg InsertBarberParams) (Barber, error)
	InsertBarberRole(ctx context.Context, arg InsertBarberRoleParams) (BarberRole, error)
	InsertBarberShop(ctx context.Context, arg InsertBarberShopParams) (BarberShop, error)
	// INSERT
	InsertBarberWithDetails(ctx context.Context, arg InsertBarberWithDetailsParams) (Barber, error)
	InsertNewOTPRequest(ctx context.Context, arg InsertNewOTPRequestParams) error
	ListBarberEmployees(ctx context.Context, arg ListBarberEmployeesParams) ([]ListBarberEmployeesRow, error)
	ListBarberShopsByBarberId(ctx context.Context, barberID uuid.UUID) ([]ListBarberShopsByBarberIdRow, error)
	ListCategoryPosition(ctx context.Context, barberShopID uuid.NullUUID) ([]ListCategoryPositionRow, error)
	ListNearbyBarberShops(ctx context.Context, arg ListNearbyBarberShopsParams) ([]ListNearbyBarberShopsRow, error)
	ListServiceCategories(ctx context.Context, barberShopID uuid.NullUUID) ([]ListServiceCategoriesRow, error)
	ListServicePackages(ctx context.Context, barberShopID uuid.UUID) ([]ViewServicePackage, error)
	ListServicesByCategory(ctx context.Context, barberShopID uuid.UUID) ([]ListServicesByCategoryRow, error)
	SearchBarberShopsByName(ctx context.Context, arg SearchBarberShopsByNameParams) ([]SearchBarberShopsByNameRow, error)
	SelectBarberRoleByIds(ctx context.Context, arg SelectBarberRoleByIdsParams) (BarberRole, error)
	// UPDATE
	UpdateBarberDetails(ctx context.Context, arg UpdateBarberDetailsParams) (Barber, error)
	UpdateBarberPassword(ctx context.Context, arg UpdateBarberPasswordParams) (Barber, error)
	UpdateBarberShopById(ctx context.Context, arg UpdateBarberShopByIdParams) (BarberShop, error)
	UpdateBarberShopChains(ctx context.Context, arg UpdateBarberShopChainsParams) (BarberShopChain, error)
	UpdateCategoryPosition(ctx context.Context, arg UpdateCategoryPositionParams) error
	UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) (Customer, error)
	UpdateOTPRequestConfirmation(ctx context.Context, arg UpdateOTPRequestConfirmationParams) error
	UpdateServiceCategory(ctx context.Context, arg UpdateServiceCategoryParams) (ServiceCategory, error)
	UpdateServiceItem(ctx context.Context, arg UpdateServiceItemParams) (ServiceItem, error)
	UpdateServicePackage(ctx context.Context, arg UpdateServicePackageParams) (ServicePackage, error)
	UpdateSessionRefreshToken(ctx context.Context, arg UpdateSessionRefreshTokenParams) error
	UpdateSessionsCustomer(ctx context.Context, arg UpdateSessionsCustomerParams) error
	UpsertServicePackageItem(ctx context.Context, arg UpsertServicePackageItemParams) error
}

var _ Querier = (*Queries)(nil)
