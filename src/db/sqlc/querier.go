// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	ChangePasswordCustomer(ctx context.Context, arg ChangePasswordCustomerParams) (Customer, error)
	CheckBarberRolePermission(ctx context.Context, arg CheckBarberRolePermissionParams) (bool, error)
	CreateAppointments(ctx context.Context, arg CreateAppointmentsParams) (CreateAppointmentsRow, error)
	CreateBarber(ctx context.Context, arg CreateBarberParams) (Barber, error)
	CreateBarberEmployee(ctx context.Context, arg CreateBarberEmployeeParams) (Barber, error)
	CreateBarberManagers(ctx context.Context, arg CreateBarberManagersParams) (BarberManager, error)
	CreateBarberRole(ctx context.Context, arg CreateBarberRoleParams) (BarberRole, error)
	CreateBarberShop(ctx context.Context, arg CreateBarberShopParams) (BarberShop, error)
	CreateBarberShopChains(ctx context.Context, arg CreateBarberShopChainsParams) (BarberShopChain, error)
	CreateBarberShopService(ctx context.Context, arg CreateBarberShopServiceParams) (BarberShopService, error)
	CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error)
	CreateServiceCategory(ctx context.Context, arg CreateServiceCategoryParams) (ServiceCategory, error)
	CreateServicesForAppointments(ctx context.Context, arg CreateServicesForAppointmentsParams) ([]BarberShopServicesAppointment, error)
	CreateSessionBarber(ctx context.Context, arg CreateSessionBarberParams) (SessionsBarber, error)
	CreateSessionsCustomer(ctx context.Context, arg CreateSessionsCustomerParams) (SessionsCustomer, error)
	DeleteBarber(ctx context.Context, id uuid.UUID) error
	DeleteBarberManagers(ctx context.Context, arg DeleteBarberManagersParams) error
	DeleteBarberRole(ctx context.Context, id uuid.UUID) error
	DeleteBarberShopChains(ctx context.Context, id uuid.UUID) error
	DeleteBarberShops(ctx context.Context, id uuid.UUID) error
	DeleteServiceCategories(ctx context.Context, id int16) error
	GetBarber(ctx context.Context, arg GetBarberParams) (GetBarberRow, error)
	GetBarberEmployee(ctx context.Context, arg GetBarberEmployeeParams) (GetBarberEmployeeRow, error)
	GetBarberRole(ctx context.Context, arg GetBarberRoleParams) (BarberRole, error)
	GetBarberShop(ctx context.Context, arg GetBarberShopParams) (GetBarberShopRow, error)
	GetBarberShopChains(ctx context.Context, id uuid.UUID) (BarberShopChain, error)
	GetCustomer(ctx context.Context, id uuid.UUID) (Customer, error)
	GetDefaultPasswordEmployee(ctx context.Context, id uuid.UUID) (sql.NullString, error)
	GetDistricts(ctx context.Context, provinceID int16) ([]District, error)
	GetProvinces(ctx context.Context) ([]Province, error)
	GetSessionBarber(ctx context.Context, id uuid.UUID) (SessionsBarber, error)
	GetSessionsCustomer(ctx context.Context, id uuid.UUID) (SessionsCustomer, error)
	GetUserBarber(ctx context.Context, arg GetUserBarberParams) (Barber, error)
	GetUserCustomer(ctx context.Context, arg GetUserCustomerParams) (Customer, error)
	GetWards(ctx context.Context, districtID int16) ([]Ward, error)
	ListAppointmentsByDate(ctx context.Context, arg ListAppointmentsByDateParams) ([]ListAppointmentsByDateRow, error)
	ListBarberManagers(ctx context.Context, barberID uuid.UUID) ([]uuid.UUID, error)
	ListBarberShops(ctx context.Context, barberID uuid.UUID) ([]ListBarberShopsRow, error)
	ListBarbersRoles(ctx context.Context, barberShopID uuid.UUID) ([]ListBarbersRolesRow, error)
	ListEmployees(ctx context.Context, arg ListEmployeesParams) ([]ListEmployeesRow, error)
	ListNearbyBarberShops(ctx context.Context, arg ListNearbyBarberShopsParams) ([]ListNearbyBarberShopsRow, error)
	ListServiceCategories(ctx context.Context, barberShopID uuid.NullUUID) ([]ServiceCategory, error)
	SearchByNameBarberShops(ctx context.Context, arg SearchByNameBarberShopsParams) ([]SearchByNameBarberShopsRow, error)
	UpdateBarber(ctx context.Context, arg UpdateBarberParams) (Barber, error)
	UpdateBarberRoles(ctx context.Context, arg UpdateBarberRolesParams) (BarberRole, error)
	UpdateBarberShop(ctx context.Context, arg UpdateBarberShopParams) (BarberShop, error)
	UpdateBarberShopChains(ctx context.Context, arg UpdateBarberShopChainsParams) (BarberShopChain, error)
	UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) (Customer, error)
	UpdateServiceCategory(ctx context.Context, arg UpdateServiceCategoryParams) error
	UpdateSessionRefreshToken(ctx context.Context, arg UpdateSessionRefreshTokenParams) error
	UpdateSessionsCustomer(ctx context.Context, arg UpdateSessionsCustomerParams) error
}

var _ Querier = (*Queries)(nil)
