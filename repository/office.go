package repository

import (
	"context"

	"github.com/sirijagadeesh/playrestapi1/models"
)

// OfficeRepo ...
type OfficeRepo interface {
	GetOfficeByCode(ctx context.Context, officeCode string) ([]*models.Office, error)
	GetOffices(ctx context.Context) ([]*models.Office, error)
}
