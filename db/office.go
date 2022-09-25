package db

import (
	"context"
	"fmt"

	"github.com/sirijagadeesh/playrestapi1/models"
)

const getOfficeByCode = `
SELECT office_code,
       city,
       phone,
       address_line1,
       address_line2,
       state,
       country,
       postal_code,
       territory
FROM offices
WHERE office_code=$1
`

// GetOfficeByCode implementation of officeRepo.
func (q *Queries) GetOfficeByCode(ctx context.Context, officeCode string) ([]*models.Office, error) {
	result := make([]*models.Office, 0)
	if err := q.db.SelectContext(ctx, &result, getOfficeByCode, officeCode); err != nil {
		return result, fmt.Errorf("failed to get data from db %w", err)
	}

	return result, nil
}

const getOffices = `
SELECT office_code,
       city,
       phone,
       address_line1,
       address_line2,
       state,
       country,
       postal_code,
       territory
FROM offices
`

// GetOffices  implementation of officeRepo.
func (q *Queries) GetOffices(ctx context.Context) ([]*models.Office, error) {
	result := make([]*models.Office, 0)
	if err := q.db.SelectContext(ctx, &result, getOffices); err != nil {
		return result, fmt.Errorf("failed to get data from db %w", err)
	}

	return result, nil
}
