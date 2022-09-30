package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	// postgresql driver.
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

const (
	port    int = 3089
	timeout int = 10
	dbTimer int = 60
)

var errMissingDBURL = errors.New("missing DB_URL in .env or as environment variable")

// env configurations passed by user.
// these read from .env or environment.
type env struct {
	DBURL string `mapstructure:"DB_URL" validate:"required,uri"`
	Port  int    `mapstructure:"PORT" validate:"required,min=3000,max=9999"`
}

// App application configuration used by application.
type App struct {
	DBConn    *sqlx.DB
	Address   string
	Timeout   time.Duration
	DBMonitor *time.Ticker
}

// Read will configs from .env or environment variables.
func Read() (*App, error) {
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("unable to read config %w", err)
	}

	readConfig := &env{
		DBURL: envString("DB_URL", ""),
		Port:  envInt("PORT", port),
	}

	if readConfig.DBURL == "" {
		return nil, errMissingDBURL
	}

	dbConn, err := sqlx.Connect("pgx", readConfig.DBURL)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to postgresql db %w", err)
	}

	return &App{
		DBConn:    dbConn,
		Address:   fmt.Sprintf(":%d", readConfig.Port),
		Timeout:   time.Second * time.Duration(timeout),
		DBMonitor: time.NewTicker(time.Second * time.Duration(dbTimer)),
	}, nil
}

func envString(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}

func envInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		value, err := strconv.Atoi(val)
		if err != nil {
			return defaultVal
		}

		return value
	}

	return defaultVal
}
