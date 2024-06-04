package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StoreMain interface {
	Querier
	DB() *pgxpool.Pool
	WithTx(tx pgx.Tx) *Queries
	ExecTx(ctx context.Context, fn func(*Queries) error) error
}

type SQLStore struct {
	db *pgxpool.Pool
	*Queries
}

func (s *SQLStore) DB() *pgxpool.Pool {
	return s.db
}

func (s *SQLStore) WithTx(tx pgx.Tx) *Queries {
	return s.Queries.WithTx(tx)
}

func (s *SQLStore) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbError := tx.Rollback(ctx); rbError != nil {
			return fmt.Errorf("tx error %v, rb error %v", err, rbError)
		}
		return err
	}
	return tx.Commit(ctx)
}

func NewStore(db *pgxpool.Pool) StoreMain {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
