package db

import (
	"context"
	"database/sql"
)

// Operation interface .. database minimum these need to be implemented.
type Operation interface {
	Close() error
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
	Stats() sql.DBStats
}

// Queries ...
type Queries struct {
	db Operation
}

// New will return Queries Object.
func New(db Operation) *Queries {
	return &Queries{db: db}
}
