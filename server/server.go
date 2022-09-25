package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirijagadeesh/playrestapi1/config"
	"github.com/sirijagadeesh/playrestapi1/handlers"
)

// Start server.
func Start(cfg *config.Config) {
	log.Printf("%#v\n", cfg.DBConn.Stats())

	defer func() {
		log.Println("closing db connection")
		log.Printf("%#v\n", cfg.DBConn.Stats())
		log.Println(cfg.DBConn.Close())
	}()

	ech := echo.New()
	ech.HideBanner = true
	ech.Use(middleware.Logger(), middleware.Recover())

	ech.GET("/ping", handlers.Ping)

	officeHandler := handlers.NewOfficeHandler(cfg.DBConn)

	ech.GET("/offices", officeHandler.GetAllOffices)
	ech.GET("/office/:office_code", officeHandler.GetOfficeByCode)

	// Start server
	go func() {
		if err := ech.Start(cfg.Address); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// err != nil && err != http.ErrServerClosed {
			ech.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)

	defer cancel()

	if err := ech.Shutdown(ctx); err != nil {
		ech.Logger.Fatal(err)
	}
}
