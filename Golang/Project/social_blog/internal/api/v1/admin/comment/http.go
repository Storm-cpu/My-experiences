package comment

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
	Create(ctx context.Context, data CreatCommentData) (*model.Comment, error)
	Update(ctx context.Context, data UpdateCommentData, userID int) (*model.Comment, error)
	View(ctx context.Context, userID int) (*model.Comment, error)
	List(ctx context.Context, lq *dbutil.ListQueryCondition, count *int64) ([]*model.Comment, error)
	Delete(ctx context.Context, userID int) error
}

func NewHTTP(svc Service, eg *echo.Group) {
	h := HTTP{svc}

	// POST /v1/admin/comments
	eg.POST("", h.create)

	// GET /v1/admin/comments/{id}
	eg.GET("/:id", h.view)

	// GET /v1/admin/comments
	eg.GET("", h.list)

	// PATCH /v1/admin/comments/{id}
	eg.PATCH("/:id", h.update)

	// DELETE /v1/admin/comments/{id}
	eg.DELETE("/:id", h.delete)
}

type ListResp struct {
	Data       []*model.Comment `json:"data"`
	TotalCount int64            `json:"total_count"`
}

type CreatCommentData struct {
	BlogID  int    `json:"blog_id" validate:"required,min=1" `
	UserID  int    `json:"user_id" validate:"required,min=1"`
	Content string `json:"content" validate:"required"`
}

type UpdateCommentData struct {
	Content string `json:"content" validate:"required"`
}

func (h *HTTP) create(c echo.Context) error {
	r := CreatCommentData{}
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

	r := UpdateCommentData{}
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
