package db

import "database/sql"

type StoreMain interface {
	Querier
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) StoreMain {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
