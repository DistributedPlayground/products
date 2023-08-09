package api

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/rs/zerolog"
)

type APIConfig struct {
	DB     *sqlx.DB
	Logger *zerolog.Logger
	Port   string
}

func heartbeat(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
