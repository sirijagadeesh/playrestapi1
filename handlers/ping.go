package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Ping ...
func Ping(ctx echo.Context) error {
	ctx.Response().Header().Set("X-Total-Count", "1")

	//nolint:wrapcheck
	return ctx.JSON(http.StatusOK,
		[]map[string]string{{"ping": fmt.Sprintf("pong at %s", time.Now().Format(time.RFC3339))}})
}
