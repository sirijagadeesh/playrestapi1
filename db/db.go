package db

import (
	"context"
	"database/sql"
)

// Operation ...
type Operation interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
}

// Queries ...
type Queries struct {
	db Operation
}

// New will return Queries Object.
func New(db Operation) *Queries {
	return &Queries{db: db}
}
