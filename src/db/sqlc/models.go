// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Appointment struct {
	AppointmentID       uuid.UUID    `json:"appointment_id"`
	CustomerID          uuid.UUID    `json:"customer_id"`
	BarberID            uuid.UUID    `json:"barber_id"`
	ServiceID           uuid.UUID    `json:"service_id"`
	AppointmentDatetime time.Time    `json:"appointment_datetime"`
	Status              int32        `json:"status"`
	CreatedAt           time.Time    `json:"created_at"`
	UpdateAt            sql.NullTime `json:"update_at"`
}

type Barber struct {
	BarberID          uuid.UUID      `json:"barber_id"`
	ShopID            uuid.NullUUID  `json:"shop_id"`
	NickName          string         `json:"nick_name"`
	FullName          string         `json:"full_name"`
	Phone             string         `json:"phone"`
	Email             string         `json:"email"`
	Gender            int32          `json:"gender"`
	Role              int32          `json:"role"`
	HashedPassword    string         `json:"hashed_password"`
	Avatar            sql.NullString `json:"avatar"`
	Status            sql.NullInt32  `json:"status"`
	PasswordChangedAt time.Time      `json:"password_changed_at"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdateAt          sql.NullTime   `json:"update_at"`
}

type BarberShop struct {
	ShopID         uuid.UUID       `json:"shop_id"`
	CodeBarberShop string          `json:"code_barber_shop"`
	OwnerID        uuid.UUID       `json:"owner_id"`
	Name           string          `json:"name"`
	Facility       int32           `json:"facility"`
	Coordinates    string          `json:"coordinates"`
	Address        string          `json:"address"`
	Image          sql.NullString  `json:"image"`
	ListImage      []string        `json:"list_image"`
	Status         int32           `json:"status"`
	Rate           sql.NullFloat64 `json:"rate"`
	IsReputation   sql.NullBool    `json:"is_reputation"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdateAt       sql.NullTime    `json:"update_at"`
}

type Customer struct {
	CustomerID        uuid.UUID      `json:"customer_id"`
	Name              string         `json:"name"`
	Email             string         `json:"email"`
	Phone             sql.NullString `json:"phone"`
	Gender            int32          `json:"gender"`
	HashedPassword    sql.NullString `json:"hashed_password"`
	Avatar            sql.NullString `json:"avatar"`
	PasswordChangedAt time.Time      `json:"password_changed_at"`
	CreateAt          time.Time      `json:"create_at"`
	UpdateAt          sql.NullTime   `json:"update_at"`
	IsSocialAuth      sql.NullBool   `json:"is_social_auth"`
}

type Service struct {
	ID          uuid.UUID       `json:"id"`
	CategoryID  uuid.UUID       `json:"category_id"`
	Name        string          `json:"name"`
	Timer       sql.NullInt32   `json:"timer"`
	Price       sql.NullFloat64 `json:"price"`
	Description sql.NullString  `json:"description"`
	Image       sql.NullString  `json:"image"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdateAt    sql.NullTime    `json:"update_at"`
}

type ServiceCategory struct {
	ID        uuid.UUID    `json:"id"`
	ShopID    uuid.UUID    `json:"shop_id"`
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdateAt  sql.NullTime `json:"update_at"`
}

type ServicesAppointment struct {
	ServicesID            uuid.UUID `json:"Services_id"`
	AppointmentsServiceID uuid.UUID `json:"Appointments_service_id"`
}

type SessionsBarber struct {
	ID           uuid.UUID `json:"id"`
	BarberID     uuid.UUID `json:"barber_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	FcmDevice    string    `json:"fcm_device"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreateAt     time.Time `json:"create_at"`
}

type SessionsCustomer struct {
	ID           uuid.UUID   `json:"id"`
	CustomerID   uuid.UUID   `json:"customer_id"`
	RefreshToken string      `json:"refresh_token"`
	UserAgent    string      `json:"user_agent"`
	ClientIp     string      `json:"client_ip"`
	FcmDevice    string      `json:"fcm_device"`
	Coordinates  interface{} `json:"coordinates"`
	IsBlocked    bool        `json:"is_blocked"`
	ExpiresAt    time.Time   `json:"expires_at"`
	CreateAt     time.Time   `json:"create_at"`
	Timezone     string      `json:"timezone"`
}
