// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: barbershops.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createBarberShop = `-- name: CreateBarberShop :one
INSERT INTO "BarberShops" (
                           barber_shop_chain_id,
                           name,
                           province_id,
                           district_id,
                           ward_id,
                           specific_location,
                           phone,
                           email,
                           website_url,
                           avatar_url,
                           cover_photo_url,
                           facade_photo_url,
                           representative_name,
                           citizen_id,
                           representative_phone,
                           representative_email,
                           representative_phone_other,
                           start_time_monday,
                           end_time_monday,
                           start_time_tuesday,
                           end_time_tuesday,
                           start_time_wednesday,
                           end_time_wednesday,
                           start_time_thursday,
                           end_time_thursday,
                           start_time_friday,
                           end_time_friday,
                           start_time_saturday,
                           end_time_saturday,
                           start_time_sunday,
                           end_time_sunday,
                           interval_scheduler,
                           coordinates
                           )
VALUES (
         $30,
$1,
$2,
$3,
$4,
$5,
$6,
$7,
         $31,

$8,
$9,
$10,
$11,
$12,
$13,
$14,
         $32,

$15,
$16,
$17,
$18,
$19,
$20,
$21,
$22,
$23,
$24,
$25,
$26,
$27,
$28,
$29,
        ST_GeographyFromText('POINT(' || $33::float8 || ' ' || $34::float8 || ')')
        ) RETURNING id, barber_shop_chain_id, name, province_id, district_id, ward_id, specific_location, phone, email, website_url, coordinates, avatar_url, cover_photo_url, facade_photo_url, representative_name, citizen_id, representative_phone, representative_email, representative_phone_other, start_time_monday, end_time_monday, start_time_tuesday, end_time_tuesday, start_time_wednesday, end_time_wednesday, start_time_thursday, end_time_thursday, start_time_friday, end_time_friday, start_time_saturday, end_time_saturday, start_time_sunday, end_time_sunday, interval_scheduler, is_main_branch, is_reputation, is_verified, default_employee_password, create_at
`

type CreateBarberShopParams struct {
	Name                     string         `json:"name"`
	ProvinceID               int16          `json:"province_id"`
	DistrictID               int16          `json:"district_id"`
	WardID                   int16          `json:"ward_id"`
	SpecificLocation         string         `json:"specific_location"`
	Phone                    string         `json:"phone"`
	Email                    string         `json:"email"`
	AvatarUrl                string         `json:"avatar_url"`
	CoverPhotoUrl            string         `json:"cover_photo_url"`
	FacadePhotoUrl           string         `json:"facade_photo_url"`
	RepresentativeName       string         `json:"representative_name"`
	CitizenID                string         `json:"citizen_id"`
	RepresentativePhone      string         `json:"representative_phone"`
	RepresentativeEmail      string         `json:"representative_email"`
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
	BarberShopChainID        uuid.NullUUID  `json:"barber_shop_chain_id"`
	WebsiteUrl               sql.NullString `json:"website_url"`
	RepresentativePhoneOther sql.NullString `json:"representative_phone_other"`
	Longitude                float64        `json:"longitude"`
	Latitude                 float64        `json:"latitude"`
}

func (q *Queries) CreateBarberShop(ctx context.Context, arg CreateBarberShopParams) (BarberShop, error) {
	row := q.db.QueryRow(ctx, createBarberShop,
		arg.Name,
		arg.ProvinceID,
		arg.DistrictID,
		arg.WardID,
		arg.SpecificLocation,
		arg.Phone,
		arg.Email,
		arg.AvatarUrl,
		arg.CoverPhotoUrl,
		arg.FacadePhotoUrl,
		arg.RepresentativeName,
		arg.CitizenID,
		arg.RepresentativePhone,
		arg.RepresentativeEmail,
		arg.StartTimeMonday,
		arg.EndTimeMonday,
		arg.StartTimeTuesday,
		arg.EndTimeTuesday,
		arg.StartTimeWednesday,
		arg.EndTimeWednesday,
		arg.StartTimeThursday,
		arg.EndTimeThursday,
		arg.StartTimeFriday,
		arg.EndTimeFriday,
		arg.StartTimeSaturday,
		arg.EndTimeSaturday,
		arg.StartTimeSunday,
		arg.EndTimeSunday,
		arg.IntervalScheduler,
		arg.BarberShopChainID,
		arg.WebsiteUrl,
		arg.RepresentativePhoneOther,
		arg.Longitude,
		arg.Latitude,
	)
	var i BarberShop
	err := row.Scan(
		&i.ID,
		&i.BarberShopChainID,
		&i.Name,
		&i.ProvinceID,
		&i.DistrictID,
		&i.WardID,
		&i.SpecificLocation,
		&i.Phone,
		&i.Email,
		&i.WebsiteUrl,
		&i.Coordinates,
		&i.AvatarUrl,
		&i.CoverPhotoUrl,
		&i.FacadePhotoUrl,
		&i.RepresentativeName,
		&i.CitizenID,
		&i.RepresentativePhone,
		&i.RepresentativeEmail,
		&i.RepresentativePhoneOther,
		&i.StartTimeMonday,
		&i.EndTimeMonday,
		&i.StartTimeTuesday,
		&i.EndTimeTuesday,
		&i.StartTimeWednesday,
		&i.EndTimeWednesday,
		&i.StartTimeThursday,
		&i.EndTimeThursday,
		&i.StartTimeFriday,
		&i.EndTimeFriday,
		&i.StartTimeSaturday,
		&i.EndTimeSaturday,
		&i.StartTimeSunday,
		&i.EndTimeSunday,
		&i.IntervalScheduler,
		&i.IsMainBranch,
		&i.IsReputation,
		&i.IsVerified,
		&i.DefaultEmployeePassword,
		&i.CreateAt,
	)
	return i, err
}

const deleteBarberShops = `-- name: DeleteBarberShops :exec
DELETE FROM "BarberShops"
WHERE id = $1
`

func (q *Queries) DeleteBarberShops(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteBarberShops, id)
	return err
}

const getBarberShop = `-- name: GetBarberShop :one
SELECT bs.id, bs.barber_shop_chain_id, bs.name, bs.province_id, bs.district_id, bs.ward_id, bs.specific_location, bs.phone, bs.email, bs.website_url, bs.coordinates, bs.avatar_url, bs.cover_photo_url, bs.facade_photo_url, bs.representative_name, bs.citizen_id, bs.representative_phone, bs.representative_email, bs.representative_phone_other, bs.start_time_monday, bs.end_time_monday, bs.start_time_tuesday, bs.end_time_tuesday, bs.start_time_wednesday, bs.end_time_wednesday, bs.start_time_thursday, bs.end_time_thursday, bs.start_time_friday, bs.end_time_friday, bs.start_time_saturday, bs.end_time_saturday, bs.start_time_sunday, bs.end_time_sunday, bs.interval_scheduler, bs.is_main_branch, bs.is_reputation, bs.is_verified, bs.default_employee_password, bs.create_at,
    COALESCE(p.name, '') AS province_name,
    COALESCE(d.name, '') AS district_name,
    COALESCE(w.name, '') AS ward_name,
    br.role_id AS role_id
FROM "BarberShops" bs
LEFT JOIN "Provinces" p ON bs."province_id" = p."id"
LEFT JOIN "Districts" d ON bs."district_id" = d."id"
LEFT JOIN "Wards" w ON bs."ward_id" = w."id"
LEFT JOIN "BarberRoles" br ON bs."id" = br."barber_shop_id"
WHERE bs.id = $1
  AND br.barber_id = $2
`

type GetBarberShopParams struct {
	ID       uuid.UUID `json:"id"`
	BarberID uuid.UUID `json:"barber_id"`
}

type GetBarberShopRow struct {
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
	ProvinceName             string         `json:"province_name"`
	DistrictName             string         `json:"district_name"`
	WardName                 string         `json:"ward_name"`
	RoleID                   pgtype.Int2    `json:"role_id"`
}

func (q *Queries) GetBarberShop(ctx context.Context, arg GetBarberShopParams) (GetBarberShopRow, error) {
	row := q.db.QueryRow(ctx, getBarberShop, arg.ID, arg.BarberID)
	var i GetBarberShopRow
	err := row.Scan(
		&i.ID,
		&i.BarberShopChainID,
		&i.Name,
		&i.ProvinceID,
		&i.DistrictID,
		&i.WardID,
		&i.SpecificLocation,
		&i.Phone,
		&i.Email,
		&i.WebsiteUrl,
		&i.Coordinates,
		&i.AvatarUrl,
		&i.CoverPhotoUrl,
		&i.FacadePhotoUrl,
		&i.RepresentativeName,
		&i.CitizenID,
		&i.RepresentativePhone,
		&i.RepresentativeEmail,
		&i.RepresentativePhoneOther,
		&i.StartTimeMonday,
		&i.EndTimeMonday,
		&i.StartTimeTuesday,
		&i.EndTimeTuesday,
		&i.StartTimeWednesday,
		&i.EndTimeWednesday,
		&i.StartTimeThursday,
		&i.EndTimeThursday,
		&i.StartTimeFriday,
		&i.EndTimeFriday,
		&i.StartTimeSaturday,
		&i.EndTimeSaturday,
		&i.StartTimeSunday,
		&i.EndTimeSunday,
		&i.IntervalScheduler,
		&i.IsMainBranch,
		&i.IsReputation,
		&i.IsVerified,
		&i.DefaultEmployeePassword,
		&i.CreateAt,
		&i.ProvinceName,
		&i.DistrictName,
		&i.WardName,
		&i.RoleID,
	)
	return i, err
}

const getDefaultPasswordEmployee = `-- name: GetDefaultPasswordEmployee :one
SELECT default_employee_password
FROM "BarberShops"
WHERE id = $1
`

func (q *Queries) GetDefaultPasswordEmployee(ctx context.Context, id uuid.UUID) (sql.NullString, error) {
	row := q.db.QueryRow(ctx, getDefaultPasswordEmployee, id)
	var default_employee_password sql.NullString
	err := row.Scan(&default_employee_password)
	return default_employee_password, err
}

const listBarberShops = `-- name: ListBarberShops :many
SELECT DISTINCT 
    bs.id,
    bs.barber_shop_chain_id,
    bs.name,
    bs.province_id,
    bs.district_id,
    bs.ward_id,
    bs.specific_location,
    bs.phone,
    bs.email,
    bs.website_url,
    bs.avatar_url,
    bs.cover_photo_url,
    bs.facade_photo_url,
    bs.is_main_branch,
    bs.is_reputation,
    COALESCE(p.name, '') AS province_name,
    COALESCE(d.name, '') AS district_name,
    COALESCE(w.name, '') AS ward_name,
    br.role_id AS role_id
FROM "BarberShops" bs
LEFT JOIN "BarberRoles" br ON bs."id" = br."barber_shop_id"
LEFT JOIN "Provinces" p ON bs."province_id" = p."id"
LEFT JOIN "Districts" d ON bs."district_id" = d."id"
LEFT JOIN "Wards" w ON bs."ward_id" = w."id"
WHERE br."barber_id" = $1
`

type ListBarberShopsRow struct {
	ID                uuid.UUID      `json:"id"`
	BarberShopChainID uuid.NullUUID  `json:"barber_shop_chain_id"`
	Name              string         `json:"name"`
	ProvinceID        int16          `json:"province_id"`
	DistrictID        int16          `json:"district_id"`
	WardID            int16          `json:"ward_id"`
	SpecificLocation  string         `json:"specific_location"`
	Phone             string         `json:"phone"`
	Email             string         `json:"email"`
	WebsiteUrl        sql.NullString `json:"website_url"`
	AvatarUrl         string         `json:"avatar_url"`
	CoverPhotoUrl     string         `json:"cover_photo_url"`
	FacadePhotoUrl    string         `json:"facade_photo_url"`
	IsMainBranch      bool           `json:"is_main_branch"`
	IsReputation      bool           `json:"is_reputation"`
	ProvinceName      string         `json:"province_name"`
	DistrictName      string         `json:"district_name"`
	WardName          string         `json:"ward_name"`
	RoleID            pgtype.Int2    `json:"role_id"`
}

func (q *Queries) ListBarberShops(ctx context.Context, barberID uuid.UUID) ([]ListBarberShopsRow, error) {
	rows, err := q.db.Query(ctx, listBarberShops, barberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListBarberShopsRow{}
	for rows.Next() {
		var i ListBarberShopsRow
		if err := rows.Scan(
			&i.ID,
			&i.BarberShopChainID,
			&i.Name,
			&i.ProvinceID,
			&i.DistrictID,
			&i.WardID,
			&i.SpecificLocation,
			&i.Phone,
			&i.Email,
			&i.WebsiteUrl,
			&i.AvatarUrl,
			&i.CoverPhotoUrl,
			&i.FacadePhotoUrl,
			&i.IsMainBranch,
			&i.IsReputation,
			&i.ProvinceName,
			&i.DistrictName,
			&i.WardName,
			&i.RoleID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listNearbyBarberShops = `-- name: ListNearbyBarberShops :many
SELECT
    id,
    barber_shop_chain_id,
    name,
    coordinates,
    "is_reputation",
    CAST(ST_X(ST_GeomFromWKB(coordinates::geometry)) AS float8) AS longitude,
    CAST(ST_Y(ST_GeomFromWKB(coordinates::geometry)) AS float8) AS latitude,
    CAST(ST_Distance(
        ST_SetSRID(ST_MakePoint($1::float, $2::float), 4326),
        coordinates::geography
    ) AS float) AS distance
FROM "BarberShops"
WHERE  ST_Distance(coordinates, ST_SetSRID(ST_MakePoint($1::float, $2::float), 4326)) <= $3::int
ORDER BY ST_Distance(coordinates, ST_SetSRID(ST_MakePoint($1::float, $2::float), 4326))
`

type ListNearbyBarberShopsParams struct {
	CurrentLongitude float64 `json:"current_longitude"`
	CurrentLatitude  float64 `json:"current_latitude"`
	DistanceInMeters int32   `json:"distance_in_meters"`
}

type ListNearbyBarberShopsRow struct {
	ID                uuid.UUID     `json:"id"`
	BarberShopChainID uuid.NullUUID `json:"barber_shop_chain_id"`
	Name              string        `json:"name"`
	Coordinates       string        `json:"coordinates"`
	IsReputation      bool          `json:"is_reputation"`
	Longitude         float64       `json:"longitude"`
	Latitude          float64       `json:"latitude"`
	Distance          float64       `json:"distance"`
}

func (q *Queries) ListNearbyBarberShops(ctx context.Context, arg ListNearbyBarberShopsParams) ([]ListNearbyBarberShopsRow, error) {
	rows, err := q.db.Query(ctx, listNearbyBarberShops, arg.CurrentLongitude, arg.CurrentLatitude, arg.DistanceInMeters)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListNearbyBarberShopsRow{}
	for rows.Next() {
		var i ListNearbyBarberShopsRow
		if err := rows.Scan(
			&i.ID,
			&i.BarberShopChainID,
			&i.Name,
			&i.Coordinates,
			&i.IsReputation,
			&i.Longitude,
			&i.Latitude,
			&i.Distance,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchByNameBarberShops = `-- name: SearchByNameBarberShops :many
SELECT
    bs.id,
    bs.barber_shop_chain_id,
    bs.name,
    bs.coordinates,
    bs."is_reputation",
    CAST(ST_X(ST_GeomFromWKB(bs.coordinates::geometry)) AS float8) AS longitude,
    CAST(ST_Y(ST_GeomFromWKB(bs.coordinates::geometry)) AS float8) AS latitude,
    CAST(ST_Distance(
        ST_SetSRID(ST_MakePoint($2::float, $3::float), 4326),
        bs.coordinates::geography
    ) AS float) AS distance
FROM "BarberShops" bs
JOIN "BarberShopChains" bsc ON bs.barber_shop_chain_id = bsc.id
WHERE bsc."name" = $1
ORDER BY ST_Distance(bs.coordinates, ST_SetSRID(ST_MakePoint($2::float, $3::float), 4326))
`

type SearchByNameBarberShopsParams struct {
	Name             string  `json:"name"`
	CurrentLongitude float64 `json:"current_longitude"`
	CurrentLatitude  float64 `json:"current_latitude"`
}

type SearchByNameBarberShopsRow struct {
	ID                uuid.UUID     `json:"id"`
	BarberShopChainID uuid.NullUUID `json:"barber_shop_chain_id"`
	Name              string        `json:"name"`
	Coordinates       string        `json:"coordinates"`
	IsReputation      bool          `json:"is_reputation"`
	Longitude         float64       `json:"longitude"`
	Latitude          float64       `json:"latitude"`
	Distance          float64       `json:"distance"`
}

func (q *Queries) SearchByNameBarberShops(ctx context.Context, arg SearchByNameBarberShopsParams) ([]SearchByNameBarberShopsRow, error) {
	rows, err := q.db.Query(ctx, searchByNameBarberShops, arg.Name, arg.CurrentLongitude, arg.CurrentLatitude)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SearchByNameBarberShopsRow{}
	for rows.Next() {
		var i SearchByNameBarberShopsRow
		if err := rows.Scan(
			&i.ID,
			&i.BarberShopChainID,
			&i.Name,
			&i.Coordinates,
			&i.IsReputation,
			&i.Longitude,
			&i.Latitude,
			&i.Distance,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBarberShop = `-- name: UpdateBarberShop :one
UPDATE "BarberShops"
SET 
    name = coalesce($2, name),
    is_main_branch = coalesce($3, is_main_branch),
    coordinates = coalesce(ST_GeographyFromText('POINT(' || $4::float8 || ' ' || $5::float8 || ')'), coordinates),
    interval_scheduler = coalesce($6, interval_scheduler),
    is_reputation = coalesce($7, is_reputation),
    is_verified = coalesce($8, is_verified),
    update_at = now()
WHERE "id" = $1
RETURNING id, barber_shop_chain_id, name, province_id, district_id, ward_id, specific_location, phone, email, website_url, coordinates, avatar_url, cover_photo_url, facade_photo_url, representative_name, citizen_id, representative_phone, representative_email, representative_phone_other, start_time_monday, end_time_monday, start_time_tuesday, end_time_tuesday, start_time_wednesday, end_time_wednesday, start_time_thursday, end_time_thursday, start_time_friday, end_time_friday, start_time_saturday, end_time_saturday, start_time_sunday, end_time_sunday, interval_scheduler, is_main_branch, is_reputation, is_verified, default_employee_password, create_at
`

type UpdateBarberShopParams struct {
	ID                uuid.UUID      `json:"id"`
	Name              sql.NullString `json:"name"`
	IsMainBranch      pgtype.Bool    `json:"is_main_branch"`
	Longitude         pgtype.Float8  `json:"longitude"`
	Latitude          pgtype.Float8  `json:"latitude"`
	IntervalScheduler pgtype.Int2    `json:"interval_scheduler"`
	IsReputation      bool           `json:"is_reputation"`
	IsVerified        bool           `json:"is_verified"`
}

func (q *Queries) UpdateBarberShop(ctx context.Context, arg UpdateBarberShopParams) (BarberShop, error) {
	row := q.db.QueryRow(ctx, updateBarberShop,
		arg.ID,
		arg.Name,
		arg.IsMainBranch,
		arg.Longitude,
		arg.Latitude,
		arg.IntervalScheduler,
		arg.IsReputation,
		arg.IsVerified,
	)
	var i BarberShop
	err := row.Scan(
		&i.ID,
		&i.BarberShopChainID,
		&i.Name,
		&i.ProvinceID,
		&i.DistrictID,
		&i.WardID,
		&i.SpecificLocation,
		&i.Phone,
		&i.Email,
		&i.WebsiteUrl,
		&i.Coordinates,
		&i.AvatarUrl,
		&i.CoverPhotoUrl,
		&i.FacadePhotoUrl,
		&i.RepresentativeName,
		&i.CitizenID,
		&i.RepresentativePhone,
		&i.RepresentativeEmail,
		&i.RepresentativePhoneOther,
		&i.StartTimeMonday,
		&i.EndTimeMonday,
		&i.StartTimeTuesday,
		&i.EndTimeTuesday,
		&i.StartTimeWednesday,
		&i.EndTimeWednesday,
		&i.StartTimeThursday,
		&i.EndTimeThursday,
		&i.StartTimeFriday,
		&i.EndTimeFriday,
		&i.StartTimeSaturday,
		&i.EndTimeSaturday,
		&i.StartTimeSunday,
		&i.EndTimeSunday,
		&i.IntervalScheduler,
		&i.IsMainBranch,
		&i.IsReputation,
		&i.IsVerified,
		&i.DefaultEmployeePassword,
		&i.CreateAt,
	)
	return i, err
}
