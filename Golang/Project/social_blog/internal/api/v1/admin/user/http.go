package user

import (
	"context"
	"net/http"
	"social_blog/internal/model"
	dbutil "social_blog/pkg/util/db"
	httputil "social_blog/pkg/util/http"

	"github.com/labstack/echo/v4"
)

type HTTP struct {
	svc Service
}

type Service interface {
	Create(ctx context.Context, data CreatUserData) (*model.User, error)
	Update(ctx context.Context, data UpdateUserData, userID int) (*model.User, error)
	View(ctx context.Context, userID int) (*model.User, error)
	List(ctx context.Context, lq *dbutil.ListQueryCondition, count *int64) ([]*model.User, error)
	Delete(ctx context.Context, userID int) error
	ChangePassword(ctx context.Context, userID int, data ChangePasswordData) error
}

func NewHTTP(svc Service, eg *echo.Group) {
	h := HTTP{svc}

	eg.POST("", h.create)
	eg.GET("/:id", h.view)
	eg.GET("", h.list)
	eg.PATCH("/:id", h.update)
	eg.DELETE("/:id", h.delete)
	eg.PATCH("/:id/password", h.changePassword)
}

type ListResp struct {
	Data       []*model.User `json:"data"`
	TotalCount int64         `json:"total_count"`
}

type CreatUserData struct {
	UserName    string `json:"username" validate:"required,min=3"`
	Password    string `json:"password" validate:"required,min=8"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	Blocked     bool   `json:"blocked"`
}

type UpdateUserData struct {
	FirstName   *string `json:"first_name,omitempty"`
	LastName    *string `json:"last_name,omitempty"`
	Email       *string `json:"email,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Blocked     *bool   `json:"blocked,omitempty"`
}

type ChangePasswordData struct {
	NewPassword        string `json:"new_password" validate:"required,min=8"`
	NewPasswordConfirm string `json:"new_password_confirm" validate:"required,eqfield=NewPassword"`
}

func (h *HTTP) create(c echo.Context) error {
	r := CreatUserData{}
	if err := c.Bind(&r); err != nil {
		return err
	}

	resp, err := h.svc.Create(c.Request().Context(), r)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) view(c echo.Context) error {
	id, err := httputil.ReqID(c)
	if err != nil {
		return err
	}

	resp, err := h.svc.View(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) list(c echo.Context) error {
	lq, err := httputil.ReqListQuery(c)
	if err != nil {
		return err
	}
	var count int64 = 0
	resp, err := h.svc.List(c.Request().Context(), lq, &count)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ListResp{resp, count})
}

func (h *HTTP) update(c echo.Context) error {
	id, err := httputil.ReqID(c)
	if err != nil {
		return err
	}

	r := UpdateUserData{}
	if err := c.Bind(&r); err != nil {
		return err
	}

	resp, err := h.svc.Update(c.Request().Context(), r, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) delete(c echo.Context) error {
	id, err := httputil.ReqID(c)
	if err != nil {
		return err
	}
	if err := h.svc.Delete(c.Request().Context(), id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func (h *HTTP) changePassword(c echo.Context) error {
	id, err := httputil.ReqID(c)
	if err != nil {
		return err
	}

	r := ChangePasswordData{}
	if err := c.Bind(&r); err != nil {
		return err
	}

	if err = c.Validate(r); err != nil {
		return err
	}

	if err := h.svc.ChangePassword(c.Request().Context(), id, r); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
