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
func Start(cfg *config.App) {
	log.Printf("%#v\n", cfg.DBConn.Stats())

	// monitor database connections.
	go func() {
		for range cfg.DBMonitor.C {
			log.Printf("%#v\n", cfg.DBConn.Stats())
		}
	}()

	defer func() {
		log.Println("closing db connection")
		log.Println(cfg.DBConn.Close())
		log.Printf("%#v\n", cfg.DBConn.Stats())
	}()

	ech := echo.New()
	ech.HideBanner = true
	ech.Use(middleware.Logger(), middleware.Recover(), middleware.CORS())

	ech.Match([]string{http.MethodHead, http.MethodGet}, "/ping", handlers.Ping)

	officeHandler := handlers.NewOfficeHandler(cfg.DBConn)

	ech.Match([]string{http.MethodHead, http.MethodGet}, "/offices", officeHandler.GetAllOffices)
	ech.Match([]string{http.MethodHead, http.MethodGet}, "/office/:office_code", officeHandler.GetOfficeByCode)

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
