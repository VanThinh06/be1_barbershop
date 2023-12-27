package db

import (
	"database/sql"
)

type StoreMain interface {
	Querier
	DB() *sql.DB
	WithTx(tx *sql.Tx) *Queries
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func (s *SQLStore) DB() *sql.DB {
	return s.db
}

func (s *SQLStore) WithTx(tx *sql.Tx) *Queries {
	return s.Queries.WithTx(tx)
}

func NewStore(db *sql.DB) StoreMain {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
