package models

// Office is struct for table offices.
//
//nolint:tagliatelle // stopped linting for json fields names.
type Office struct {
	OfficeCode   string  `json:"office_code" db:"office_code" validate:"required,max=10"`
	City         string  `json:"city" db:"city" validate:"required,max=50"`
	Phone        string  `json:"phone" db:"phone" validate:"required,max=50"`
	AddressLine1 string  `json:"address_line_1" db:"address_line1" validate:"required,max=50"`
	AddressLine2 *string `json:"address_line_2,string" db:"address_line2" validate:"max=50"`
	State        *string `json:"state,string" db:"state" validate:"max=50"`
	Country      string  `json:"country" db:"country" validate:"max=50"`
	PostalCode   string  `json:"postal_code" db:"postal_code" validate:"max=15"`
	Territory    string  `json:"territory" db:"territory" validate:"max=10"`
}
