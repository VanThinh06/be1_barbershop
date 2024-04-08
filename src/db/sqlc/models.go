// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Appointment struct {
	ID                  uuid.UUID `json:"id"`
	BarberShopID        uuid.UUID `json:"barber_shop_id"`
	CustomerID          uuid.UUID `json:"customer_id"`
	BarberID            uuid.UUID `json:"barber_id"`
	AppointmentDateTime time.Time `json:"appointment_date_time"`
	Status              int16     `json:"status"`
	CreateAt            time.Time `json:"create_at"`
	UpdateAt            time.Time `json:"update_at"`
}

type Barber struct {
	ID                uuid.UUID      `json:"id"`
	GenderID          int16          `json:"gender_id"`
	Phone             string         `json:"phone"`
	NickName          string         `json:"nick_name"`
	Email             sql.NullString `json:"email"`
	HashedPassword    sql.NullString `json:"hashed_password"`
	FullName          string         `json:"full_name"`
	Haircut           bool           `json:"haircut"`
	AvatarUrl         sql.NullString `json:"avatar_url"`
	PasswordChangedAt time.Time      `json:"password_changed_at"`
	CreateAt          time.Time      `json:"create_at"`
}

type BarberManager struct {
	BarberID  uuid.UUID `json:"barber_id"`
	ManagerID uuid.UUID `json:"manager_id"`
}

type BarberReview struct {
	ID         uuid.UUID      `json:"id"`
	BarberID   uuid.UUID      `json:"barber_id"`
	CustomerID uuid.UUID      `json:"customer_id"`
	Rating     int16          `json:"rating"`
	Comment    sql.NullString `json:"comment"`
	CreateAt   time.Time      `json:"create_at"`
	UpdateAt   time.Time      `json:"update_at"`
}

type BarberRole struct {
	ID           uuid.UUID `json:"id"`
	BarberID     uuid.UUID `json:"barber_id"`
	BarberShopID uuid.UUID `json:"barber_shop_id"`
	RoleID       int16     `json:"role_id"`
}

type BarberShop struct {
	ID                       uuid.UUID      `json:"id"`
	BarberShopChainID        uuid.NullUUID  `json:"barber_shop_chain_id"`
	Name                     string         `json:"name"`
	ProvinceID               int16          `json:"province_id"`
	DistrictID               int16          `json:"district_id"`
	WardID                   int16          `json:"ward_id"`
	SpecificLocation         string         `json:"specific_location"`
	Phone                    string         `json:"phone"`
	Email                    string         `json:"email"`
	WebsiteUrl               sql.NullString `json:"website_url"`
	Coordinates              string         `json:"coordinates"`
	AvatarUrl                string         `json:"avatar_url"`
	CoverPhotoUrl            string         `json:"cover_photo_url"`
	FacadePhotoUrl           string         `json:"facade_photo_url"`
	RepresentativeName       string         `json:"representative_name"`
	CitizenID                string         `json:"citizen_id"`
	RepresentativePhone      string         `json:"representative_phone"`
	RepresentativeEmail      string         `json:"representative_email"`
	RepresentativePhoneOther sql.NullString `json:"representative_phone_other"`
	StartTimeMonday          pgtype.Time    `json:"start_time_monday"`
	EndTimeMonday            pgtype.Time    `json:"end_time_monday"`
	StartTimeTuesday         pgtype.Time    `json:"start_time_tuesday"`
	EndTimeTuesday           pgtype.Time    `json:"end_time_tuesday"`
	StartTimeWednesday       pgtype.Time    `json:"start_time_wednesday"`
	EndTimeWednesday         pgtype.Time    `json:"end_time_wednesday"`
	StartTimeThursday        pgtype.Time    `json:"start_time_thursday"`
	EndTimeThursday          pgtype.Time    `json:"end_time_thursday"`
	StartTimeFriday          pgtype.Time    `json:"start_time_friday"`
	EndTimeFriday            pgtype.Time    `json:"end_time_friday"`
	StartTimeSaturday        pgtype.Time    `json:"start_time_saturday"`
	EndTimeSaturday          pgtype.Time    `json:"end_time_saturday"`
	StartTimeSunday          pgtype.Time    `json:"start_time_sunday"`
	EndTimeSunday            pgtype.Time    `json:"end_time_sunday"`
	IntervalScheduler        int16          `json:"interval_scheduler"`
	IsMainBranch             bool           `json:"is_main_branch"`
	IsReputation             bool           `json:"is_reputation"`
	IsVerified               bool           `json:"is_verified"`
	DefaultEmployeePassword  sql.NullString `json:"default_employee_password"`
	CreateAt                 time.Time      `json:"create_at"`
}

type BarberShopChain struct {
	ID              uuid.UUID      `json:"id"`
	Name            string         `json:"name"`
	ChainIdentifier string         `json:"chain_identifier"`
	Founder         string         `json:"founder"`
	FoundingDate    pgtype.Date    `json:"founding_date"`
	Website         sql.NullString `json:"website"`
}

type BarberShopReview struct {
	ID           uuid.UUID      `json:"id"`
	CustomerID   uuid.UUID      `json:"customer_id"`
	BarberShopID uuid.UUID      `json:"barber_shop_id"`
	Rating       int16          `json:"rating"`
	Comment      sql.NullString `json:"comment"`
	CreateAt     time.Time      `json:"create_at"`
	UpdateAt     time.Time      `json:"update_at"`
}

type BarberShopService struct {
	ID                   uuid.UUID      `json:"id"`
	BarbershopCategoryID uuid.UUID      `json:"barbershop_category_id"`
	GenderID             int16          `json:"gender_id"`
	Name                 string         `json:"name"`
	Timer                int16          `json:"timer"`
	Price                float32        `json:"price"`
	Description          sql.NullString `json:"description"`
	ImageUrl             sql.NullString `json:"image_url"`
}

type BarberShopServiceCategory struct {
	ID                uuid.UUID `json:"id"`
	BarberShopID      uuid.UUID `json:"barber_shop_id"`
	ServiceCategoryID int16     `json:"service_category_id"`
}

type BarberShopServicesAppointment struct {
	BarberShopServicesID  uuid.UUID `json:"BarberShopServices_id"`
	AppointmentsServiceID uuid.UUID `json:"AppointmentsService_id"`
}

type Customer struct {
	ID                uuid.UUID      `json:"id"`
	Name              string         `json:"name"`
	Email             string         `json:"email"`
	Phone             sql.NullString `json:"phone"`
	GenderID          int16          `json:"gender_id"`
	HashedPassword    sql.NullString `json:"hashed_password"`
	Avatar            sql.NullString `json:"avatar"`
	IsSocialAuth      pgtype.Bool    `json:"is_social_auth"`
	PasswordChangedAt time.Time      `json:"password_changed_at"`
	CreateAt          time.Time      `json:"create_at"`
}

type District struct {
	ID         int16  `json:"id"`
	Name       string `json:"name"`
	ProvinceID int16  `json:"province_id"`
}

type Gender struct {
	ID   int16  `json:"id"`
	Name string `json:"name"`
}

type Province struct {
	ID   int16  `json:"id"`
	Name string `json:"name"`
}

type Role struct {
	ID   int16          `json:"id"`
	Name string         `json:"name"`
	Type sql.NullString `json:"type"`
}

type ServiceCategory struct {
	ID       int16  `json:"id"`
	Name     string `json:"name"`
	IsGlobal bool   `json:"is_global"`
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
}

type Ward struct {
	ID         int16  `json:"id"`
	Name       string `json:"name"`
	DistrictID int16  `json:"district_id"`
}
