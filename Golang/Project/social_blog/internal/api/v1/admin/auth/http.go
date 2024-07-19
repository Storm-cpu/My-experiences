package auth

import (
	"context"
	"net/http"
	"social_blog/internal/model"

	"github.com/labstack/echo/v4"
)

type HTTP struct {
	svc Service
}

type Service interface {
	Authenticate(ctx context.Context, data Credentials) (*model.AuthToken, error)
	RefreshToken(ctx context.Context, data RefreshTokenData) (*model.AuthToken, error)
}

func NewHTTP(svc Service, eg *echo.Group) {
	h := HTTP{svc}

	// POST v1/admin/login
	eg.POST("/login", h.login)

	// POST v1/admin/refresh-token
	eg.POST("/refresh-token", h.refreshToken)
}

// Credentials contains login request data
type Credentials struct {
	Username string `json:"username" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
}

type RefreshTokenData struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
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

func (h *HTTP) refreshToken(c echo.Context) error {
	r := RefreshTokenData{}
	if err := c.Bind(&r); err != nil {
		return err
	}
	resp, err := h.svc.RefreshToken(c.Request().Context(), r)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
