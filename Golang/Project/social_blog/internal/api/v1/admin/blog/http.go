package blog

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
	Create(ctx context.Context, data CreateBlogData) (*model.Blog, error)
	View(ctx context.Context, blogID int) (*model.Blog, error)
	List(ctx context.Context, lq *dbutil.ListQueryCondition, count *int64) ([]*model.Blog, error)
	Update(ctx context.Context, data UpdateBlogData, blogID int) (*model.Blog, error)
	Delete(ctx context.Context, blogID int) error
}

func NewHTTP(svc Service, eg *echo.Group) {
	h := HTTP{svc}

	// POST /v1/admin/blogs
	eg.POST("", h.create)

	// GET /v1/admin/blogs/{id}
	eg.GET("/:id", h.view)

	// GET /v1/admin/blogs
	eg.GET("", h.list)

	// PATCH /v1/admin/blogs/{id}
	eg.PATCH("/:id", h.update)

	// DELETE /v1/admin/blogs/{id}
	eg.DELETE("/:id", h.delete)
}

// ListResp contains list of blogs and current page number response
type ListResp struct {
	Data       []*model.Blog `json:"data"`
	TotalCount int64         `json:"total_count"`
}

// CreateBlogData contains user data from request
type CreateBlogData struct {
	Title      string `json:"title" validate:"required"`
	Content    string `json:"content" validate:"required"`
	AuthorID   int    `json:"author_id" gorm:"not null" validate:"required,min=1"`
	Status     string `json:"status" validate:"max=50"`
	Visibility string `json:"visibility"  validate:"max=50"`
}

// UpdateBlogData contains user data from request
type UpdateBlogData struct {
	Title      *string `json:"title,omitempty"`
	Content    *string `json:"content,omitempty"`
	AuthorID   *int    `json:"author_id,omitempty" gorm:"not null"`
	Status     *string `json:"status,omitempty" gorm:"not null"`
	Visibility *string `json:"visibility,omitempty" gorm:"not null"`
}

func (h *HTTP) create(c echo.Context) error {
	r := CreateBlogData{}
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

	r := UpdateBlogData{}
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
