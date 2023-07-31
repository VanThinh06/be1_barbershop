// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package db

import (
	"time"

	"github.com/google/uuid"
	null "gopkg.in/guregu/null.v4"
)

type Barber struct {
	Username          string        `json:"username"`
	FullName          string        `json:"full_name"`
	Email             string        `json:"email"`
	HashedPassword    string        `json:"hashed_password"`
	Avatar            null.String   `json:"avatar"`
	Role              null.String   `json:"role"`
	Status            null.String   `json:"status"`
	StoreID           uuid.NullUUID `json:"store_id"`
	StoreManager      []uuid.UUID   `json:"store_manager"`
	PasswordChangedAt time.Time     `json:"password_changed_at"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdateAt          null.Time     `json:"update_at"`
}

type Service struct {
	ID          uuid.UUID   `json:"id"`
	StoreID     uuid.UUID   `json:"store_id"`
	Work        string      `json:"work"`
	Timer       int32       `json:"timer"`
	Price       float32     `json:"price"`
	Description null.String `json:"description"`
	Image       null.String `json:"image"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdateAt    null.Time   `json:"update_at"`
}

type SessionsBarber struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	IsManager    bool      `json:"is_manager"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	FcmDevice    string    `json:"fcm_device"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreateAt     time.Time `json:"create_at"`
}

type Store struct {
	ID         uuid.UUID   `json:"id"`
	NameID     string      `json:"name_id"`
	NameStore  string      `json:"name_store"`
	Location   float32     `json:"location"`
	Address    string      `json:"address"`
	Image      null.String `json:"image"`
	ListImage  []string    `json:"list_image"`
	ManagerID  []uuid.UUID `json:"manager_id"`
	EmployeeID []uuid.UUID `json:"employee_id"`
	Status     int32       `json:"status"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdateAt   null.Time   `json:"update_at"`
}
