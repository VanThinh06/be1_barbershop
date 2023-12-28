package db

import (
	"context"
	"database/sql"
	"fmt"
)

type StoreMain interface {
	Querier
	DB() *sql.DB
	WithTx(tx *sql.Tx) *Queries
	ExecTx(ctx context.Context, fn func(*Queries) error) error
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

func (s *SQLStore) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbError := tx.Rollback(); rbError != nil {
		return fmt.Errorf("tx error %v, rb error %v", err, rbError)
		}
		return err
	}
	return tx.Commit()
}


func NewStore(db *sql.DB) StoreMain {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
