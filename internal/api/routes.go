package api

import (
	"consult/internal/api/handler"
	"consult/internal/api/middleware"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (c *CustomValidator) Validate(i interface{}) error {
	if err := c.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (s *Server) routes(h *handler.Handler, m *middleware.Middleware) {
	m.Config(s.app, []string{"/login"})

	s.app.Validator = &CustomValidator{validator: validator.New()}

	s.app.POST("/login", h.Login)
}
