package category

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
	Create(ctx context.Context, data model.Category) (*model.Category, error)
	View(ctx context.Context, categoryID int) (*model.Category, error)
	List(ctx context.Context, lq *dbutil.ListQueryCondition, count *int64) ([]*model.Category, error)
	Update(ctx context.Context, data model.Category, categoryID int) (*model.Category, error)
	Delete(ctx context.Context, categoryID int) error
}

func NewHTTP(svc Service, eg *echo.Group) {
	h := HTTP{svc}

	// POST /v1/admin/categories
	eg.POST("", h.create)

	// GET /v1/admin/categories/{id}
	eg.GET("/:id", h.view)

	// GET /v1/admin/categories
	eg.GET("", h.list)

	// PATCH /v1/admin/categories/{id}
	eg.PATCH("/:id", h.update)

	// DELETE /v1/admin/categories/{id}
	eg.DELETE("/:id", h.delete)
}

// ListResp contains list of categories and current page number response
type ListResp struct {
	Data       []*model.Category `json:"data"`
	TotalCount int64             `json:"total_count"`
}

func (h *HTTP) create(c echo.Context) error {
	r := model.Category{}
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

	r := model.Category{}
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
