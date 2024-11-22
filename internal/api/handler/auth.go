package handler

import (
	"consult/internal/api/repository"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func (h *Handler) Login(c echo.Context) error {
	u := new(struct {
		Email    string
		Password string
	})

	if err := c.Bind(u); err != nil {
		return h.BadRequest(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(u); err != nil {
		return err
	}

	user, err := h.repository.GetUserByEmail(u.Email)

	claims := &JwtClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
		},
	}

	if err != nil || user.Password != u.Password {
		return h.BadRequest(http.StatusBadRequest, "Credenciais inválidas.")
	}

	jwtSecret := []byte(h.config.GetString("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(jwtSecret)

	if err != nil {
		return h.BadRequest(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func (h *Handler) Register(c echo.Context) error {
	u := new(struct {
		Email    string
		Password string
		Role     string
	})

	if err := c.Bind(u); err != nil {
		return h.BadRequest(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(u); err != nil {
		return err
	}

	if user, _ := h.repository.GetUserByEmail(u.Email); user.Email != "" {
		return h.BadRequest(http.StatusConflict, "Já existe um email com essa conta.")
	}

	user, err := h.repository.SaveUser(repository.User{Email: u.Email, Password: u.Password, Role: u.Role})

	if err != nil {
		return h.BadRequest(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"id":    user.Id,
		"email": user.Email,
		"role":  user.Role,
	})
}
