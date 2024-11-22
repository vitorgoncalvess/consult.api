package handler

import (
	"consult/internal/api/database"
	"consult/internal/api/repository"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Handler struct {
	repository *repository.Repository
	config     *viper.Viper
}

func New(database *database.Database, config *viper.Viper) *Handler {
	repository := repository.New(database, config)
	return &Handler{repository, config}
}

func (h *Handler) BadRequest(status int, message string) *echo.HTTPError {
	return echo.NewHTTPError(status, message)
}
