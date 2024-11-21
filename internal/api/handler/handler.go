package handler

import (
	"consult/internal/api/database"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Handler struct {
	database *database.Database
	config   *viper.Viper
}

func New(database *database.Database, config *viper.Viper) *Handler {
	return &Handler{database, config}
}

func (h *Handler) BadRequest(status int, message string) *echo.HTTPError {
	return echo.NewHTTPError(status, message)
}
