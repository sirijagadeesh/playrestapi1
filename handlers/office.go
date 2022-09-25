package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirijagadeesh/playrestapi1/db"
	"github.com/sirijagadeesh/playrestapi1/repository"
)

// OfficeHandler ...
type OfficeHandler struct {
	office repository.OfficeRepo
}

// NewOfficeHandler ...
func NewOfficeHandler(dbOp db.Operation) *OfficeHandler {
	return &OfficeHandler{office: db.New(dbOp)}
}

// GetOfficeByCode echo handler implementation to get Offices by Code.
func (oh *OfficeHandler) GetOfficeByCode(ctx echo.Context) error {
	OfficeID := ctx.Param("office_code")

	data, err := oh.office.GetOfficeByCode(ctx.Request().Context(), OfficeID)
	if err != nil {
		log.Println(err)

		return ctx.JSON(http.StatusInternalServerError, []map[string]string{{"error": "Internal Error"}}) //nolint:wrapcheck
	}

	ctx.Response().Header().Set("X-Total-Count", strconv.Itoa(len(data)))

	return ctx.JSON(http.StatusOK, data) //nolint:wrapcheck
}

// GetAllOffices echo handler to get all offices.
func (oh *OfficeHandler) GetAllOffices(ctx echo.Context) error {
	data, err := oh.office.GetOffices(ctx.Request().Context())
	if err != nil {
		log.Println(err)

		return ctx.JSON(http.StatusInternalServerError, []map[string]string{{"error": "Internal Error"}}) //nolint:wrapcheck
	}

	ctx.Response().Header().Set("X-Total-Count", strconv.Itoa(len(data)))

	return ctx.JSON(http.StatusOK, data) //nolint:wrapcheck
}
