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

type Barber struct {
	ID                uuid.UUID          `json:"id"`
	GenderID          pgtype.Int2        `json:"gender_id"`
	Phone             string             `json:"phone"`
	NickName          string             `json:"nick_name"`
	Email             string             `json:"email"`
	HashedPassword    sql.NullString     `json:"hashed_password"`
	FullName          string             `json:"full_name"`
	Haircut           bool               `json:"haircut"`
	WorkStatus        bool               `json:"work_status"`
	AvatarUrl         sql.NullString     `json:"avatar_url"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	CreateAt          time.Time          `json:"create_at"`
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

type CategoryPosition struct {
	BarberShopID uuid.UUID `json:"barber_shop_id"`
	CategoryID   int16     `json:"category_id"`
	Position     int16     `json:"position"`
	Visible      bool      `json:"visible"`
}

type Customer struct {
	ID                uuid.UUID          `json:"id"`
	Name              string             `json:"name"`
	Email             string             `json:"email"`
	Phone             sql.NullString     `json:"phone"`
	GenderID          pgtype.Int2        `json:"gender_id"`
	HashedPassword    sql.NullString     `json:"hashed_password"`
	Avatar            sql.NullString     `json:"avatar"`
	IsSocialAuth      pgtype.Bool        `json:"is_social_auth"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	CreateAt          time.Time          `json:"create_at"`
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

type OTPRequest struct {
	ID          uuid.UUID `json:"id"`
	BarberID    uuid.UUID `json:"barber_id"`
	Otp         string    `json:"otp"`
	RequestedAt time.Time `json:"requested_at"`
	IsConfirm   bool      `json:"is_confirm"`
}

type OTPRequestsCustomer struct {
	ID          uuid.UUID `json:"id"`
	CustomerID  uuid.UUID `json:"customer_id"`
	Otp         string    `json:"otp"`
	RequestedAt time.Time `json:"requested_at"`
	IsConfirm   bool      `json:"is_confirm"`
}

type Permission struct {
	ID          int16          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}

type Province struct {
	ID   int16  `json:"id"`
	Name string `json:"name"`
}

type Role struct {
	ID   int16  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type RolePermission struct {
	ID           int16 `json:"id"`
	RoleID       int16 `json:"role_id"`
	PermissionID int16 `json:"permission_id"`
}

type ServiceCategory struct {
	ID           int16         `json:"id"`
	Name         string        `json:"name"`
	BarberShopID uuid.NullUUID `json:"barber_shop_id"`
}

type ServiceItem struct {
	ID                uuid.UUID        `json:"id"`
	BarberShopID      uuid.UUID        `json:"barber_shop_id"`
	CategoryID        int16            `json:"category_id"`
	GenderID          int16            `json:"gender_id"`
	Name              string           `json:"name"`
	Timer             int16            `json:"timer"`
	Price             float32          `json:"price"`
	DiscountPrice     pgtype.Float4    `json:"discount_price"`
	DiscountStartTime pgtype.Timestamp `json:"discount_start_time"`
	DiscountEndTime   pgtype.Timestamp `json:"discount_end_time"`
	Description       sql.NullString   `json:"description"`
	ImageUrl          sql.NullString   `json:"image_url"`
	IsActive          bool             `json:"is_active"`
}

type ServicePackage struct {
	ID                uuid.UUID        `json:"id"`
	BarberShopID      uuid.UUID        `json:"barber_shop_id"`
	Name              string           `json:"name"`
	GenderID          int16            `json:"gender_id"`
	Timer             int16            `json:"timer"`
	Price             float32          `json:"price"`
	DiscountPrice     pgtype.Float4    `json:"discount_price"`
	DiscountStartTime pgtype.Timestamp `json:"discount_start_time"`
	DiscountEndTime   pgtype.Timestamp `json:"discount_end_time"`
	Description       sql.NullString   `json:"description"`
	ImageUrl          sql.NullString   `json:"image_url"`
	IsActive          bool             `json:"is_active"`
}

type ServicePackageItem struct {
	ID               uuid.UUID `json:"id"`
	ServicePackageID uuid.UUID `json:"service_package_id"`
	ServiceItemID    uuid.UUID `json:"service_item_id"`
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

type ViewServicePackage struct {
	ID                            uuid.UUID        `json:"id"`
	BarberShopID                  uuid.UUID        `json:"barber_shop_id"`
	ComboServiceGender            int16            `json:"combo_service_gender"`
	ComboServiceName              string           `json:"combo_service_name"`
	ComboServiceTimer             int16            `json:"combo_service_timer"`
	ComboServicePrice             float32          `json:"combo_service_price"`
	ComboServiceDiscountPrice     pgtype.Float4    `json:"combo_service_discount_price"`
	ComboServiceDiscountStartTime pgtype.Timestamp `json:"combo_service_discount_start_time"`
	ComboServiceDiscountEndTime   pgtype.Timestamp `json:"combo_service_discount_end_time"`
	ComboServiceDescription       sql.NullString   `json:"combo_service_description"`
	ComboServiceImageUrl          sql.NullString   `json:"combo_service_image_url"`
	ComboServiceIsActive          bool             `json:"combo_service_is_active"`
	ServiceItems                  []ServiceItem    `json:"service_items"`
}

type Ward struct {
	ID         int16  `json:"id"`
	Name       string `json:"name"`
	DistrictID int16  `json:"district_id"`
}
