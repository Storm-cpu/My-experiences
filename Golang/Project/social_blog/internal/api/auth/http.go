package auth

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HTTP struct {
	svc Service
}

type Service interface {
	Authenticate(ctx context.Context, data Credentials) (*RepMessage, error)
}

func NewHTTP(svc Service, e *echo.Echo) {
	h := HTTP{svc}

	e.POST("/login", h.login)
}

// Credentials contains login request data
type Credentials struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RepMessage struct {
	Message string `json:"message"`
}

func (h *HTTP) login(c echo.Context) error {
	r := Credentials{}
	if err := c.Bind(&r); err != nil {
		return err
	}
	resp, err := h.svc.Authenticate(c.Request().Context(), r)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
