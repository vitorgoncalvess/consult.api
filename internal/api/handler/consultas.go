package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetConsultasByUser(c echo.Context) error {
	return c.String(http.StatusOK, "vai tomando")
}
