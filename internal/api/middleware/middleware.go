package middleware

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Middleware struct {
	jwtSecret []byte
}

func New(config *viper.Viper) *Middleware {
	jwtSecret := []byte(config.GetString("JWT_SECRET"))

	return &Middleware{
		jwtSecret: jwtSecret,
	}
}

func (m *Middleware) Config(e *echo.Echo, publicRoutes []string) {
	e.Use(echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			for _, route := range publicRoutes {
				if route == c.Request().URL.Path {
					return true
				}
			}
			return false
		},
		SigningKey: m.jwtSecret,
	}))
}
