package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	queryObj := New(tx)
	err = fn(queryObj)
	if err != nil {
		if rollback_err := tx.Rollback(); rollback_err != nil {
			return fmt.Errorf("Tx Error: %v, Rollback Error: %v", err, rollback_err)
		}

	}
	return tx.Commit()
}
