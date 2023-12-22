// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: barber_shops.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const barberShopInChain = `-- name: BarberShopInChain :one
SELECT EXISTS (
  SELECT 1
  FROM "BarberShops"
  WHERE "owner_id" = $1
    AND "chain_id" IS NULL
) AS "barbershop_not_in_chain"
`

func (q *Queries) BarberShopInChain(ctx context.Context, ownerID uuid.UUID) (bool, error) {
	row := q.db.QueryRowContext(ctx, barberShopInChain, ownerID)
	var barbershop_not_in_chain bool
	err := row.Scan(&barbershop_not_in_chain)
	return barbershop_not_in_chain, err
}

const createBarberShop = `-- name: CreateBarberShop :one

INSERT INTO "BarberShops" (
                           owner_id,
                           chain_id,
                           name,
                           coordinates,
                           address,
                           image,
                           facility
                           )
VALUES (
        $1,
        $2,
        $3,
        ST_GeographyFromText('POINT(' || $4::float8 || ' ' || $5::float8 || ')'),
        $6,
        $7,
        $8
        ) RETURNING shop_id, owner_id, chain_id, name, facility, address, image, list_image, status, coordinates, rate, is_reputation, created_at, updated_at
`

type CreateBarberShopParams struct {
	OwnerID   uuid.UUID      `json:"owner_id"`
	ChainID   uuid.NullUUID  `json:"chain_id"`
	Name      string         `json:"name"`
	Longitude float64        `json:"longitude"`
	Latitude  float64        `json:"latitude"`
	Address   string         `json:"address"`
	Image     sql.NullString `json:"image"`
	Facility  int32          `json:"facility"`
}

func (q *Queries) CreateBarberShop(ctx context.Context, arg CreateBarberShopParams) (BarberShop, error) {
	row := q.db.QueryRowContext(ctx, createBarberShop,
		arg.OwnerID,
		arg.ChainID,
		arg.Name,
		arg.Longitude,
		arg.Latitude,
		arg.Address,
		arg.Image,
		arg.Facility,
	)
	var i BarberShop
	err := row.Scan(
		&i.ShopID,
		&i.OwnerID,
		&i.ChainID,
		&i.Name,
		&i.Facility,
		&i.Address,
		&i.Image,
		pq.Array(&i.ListImage),
		&i.Status,
		&i.Coordinates,
		&i.Rate,
		&i.IsReputation,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findBarberShopsNearbyLocations = `-- name: FindBarberShopsNearbyLocations :many

SELECT
    owner_id,
    shop_id,
    status,
    name,
    coordinates,
    address,
    image,
    list_image,
    created_at,
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

type FindBarberShopsNearbyLocationsParams struct {
	CurrentLongitude float64 `json:"current_longitude"`
	CurrentLatitude  float64 `json:"current_latitude"`
	DistanceInMeters int32   `json:"distance_in_meters"`
}

type FindBarberShopsNearbyLocationsRow struct {
	OwnerID     uuid.UUID      `json:"owner_id"`
	ShopID      uuid.UUID      `json:"shop_id"`
	Status      int32          `json:"status"`
	Name        string         `json:"name"`
	Coordinates string         `json:"coordinates"`
	Address     string         `json:"address"`
	Image       sql.NullString `json:"image"`
	ListImage   []string       `json:"list_image"`
	CreatedAt   time.Time      `json:"created_at"`
	Longitude   float64        `json:"longitude"`
	Latitude    float64        `json:"latitude"`
	Distance    float64        `json:"distance"`
}

func (q *Queries) FindBarberShopsNearbyLocations(ctx context.Context, arg FindBarberShopsNearbyLocationsParams) ([]FindBarberShopsNearbyLocationsRow, error) {
	rows, err := q.db.QueryContext(ctx, findBarberShopsNearbyLocations, arg.CurrentLongitude, arg.CurrentLatitude, arg.DistanceInMeters)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []FindBarberShopsNearbyLocationsRow{}
	for rows.Next() {
		var i FindBarberShopsNearbyLocationsRow
		if err := rows.Scan(
			&i.OwnerID,
			&i.ShopID,
			&i.Status,
			&i.Name,
			&i.Coordinates,
			&i.Address,
			&i.Image,
			pq.Array(&i.ListImage),
			&i.CreatedAt,
			&i.Longitude,
			&i.Latitude,
			&i.Distance,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBarberShop = `-- name: GetBarberShop :one

SELECT "shop_id"
FROM "BarberShops"
WHERE shop_id = $1
LIMIT 1
`

func (q *Queries) GetBarberShop(ctx context.Context, shopID uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRowContext(ctx, getBarberShop, shopID)
	var shop_id uuid.UUID
	err := row.Scan(&shop_id)
	return shop_id, err
}

const updateChainForBarberShops = `-- name: UpdateChainForBarberShops :exec
UPDATE "BarberShops"
SET "chain_id" = $2::uuid
WHERE "owner_id" = $1 AND "chain_id" IS NULL
`

type UpdateChainForBarberShopsParams struct {
	OwnerID uuid.UUID `json:"owner_id"`
	ChainID uuid.UUID `json:"chain_id"`
}

func (q *Queries) UpdateChainForBarberShops(ctx context.Context, arg UpdateChainForBarberShopsParams) error {
	_, err := q.db.ExecContext(ctx, updateChainForBarberShops, arg.OwnerID, arg.ChainID)
	return err
}