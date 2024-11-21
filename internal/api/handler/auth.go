package handler

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type JwtClaims struct {
	Id    int
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (h *Handler) Login(c echo.Context) error {
	u := new(User)

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(u); err != nil {
		return err
	}

	claims := &JwtClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}

	res := h.conn.QueryRow("SELECT * FROM user WHERE email = ?", u.Email)

	var password string
	err := res.Scan(&claims.Id, &claims.Email, &password)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if password != u.Password {
		return echo.NewHTTPError(http.StatusBadRequest, "Credenciais invalidas")
	}

	jwtSecret := []byte(h.config.GetString("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(jwtSecret)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
