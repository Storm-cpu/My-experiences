package category

import (
	"context"
	"net/http"
	"social_blog/internal/model"
	"social_blog/pkg/server"
	dbutil "social_blog/pkg/util/db"
	structutil "social_blog/pkg/util/struct"
)

var (
	ErrCategoryNotFound = server.NewHTTPError(http.StatusBadRequest, "CATEGORY_NOTFOUND", "Category không tồn tại")
)

// Create creates a new Category
func (c *Category) Create(ctx context.Context, data model.Category) (*model.Category, error) {
	rec := &model.Category{
		Name: data.Name,
	}

	if err := c.ctdb.Create(c.db.WithContext(ctx), rec); err != nil {
		return nil, server.NewHTTPInternalError("Error creating blog").SetInternal(err)
	}

	return rec, nil
}

// View returns a single Category
func (c *Category) View(ctx context.Context, categoryID int) (*model.Category, error) {
	rec := new(model.Category)
	if err := c.ctdb.View(c.db.WithContext(ctx), rec, categoryID); err != nil {
		return nil, ErrCategoryNotFound.SetInternal(err)
	}
	return rec, nil
}

// List returns a list of Blogs
func (c *Category) List(ctx context.Context, lq *dbutil.ListQueryCondition, count *int64) ([]*model.Category, error) {
	var data []*model.Category
	if err := c.ctdb.List(c.db.WithContext(ctx), &data, lq, count); err != nil {
		return nil, server.NewHTTPInternalError("Error listing blog").SetInternal(err)
	}
	return data, nil
}

// Update updates Category information
func (c *Category) Update(ctx context.Context, data model.Category, categoryID int) (*model.Category, error) {
	update := structutil.ToMap(data)
	if err := c.ctdb.Update(c.db.WithContext(ctx), update, categoryID); err != nil {
		return nil, server.NewHTTPInternalError("Error updating blog").SetInternal(err)
	}

	rec := new(model.Category)
	if err := c.ctdb.View(c.db.WithContext(ctx), rec, categoryID); err != nil {
		return nil, ErrCategoryNotFound.SetInternal(err)
	}

	return rec, nil
}

// Delete deletes a Category
func (c *Category) Delete(ctx context.Context, categoryID int) error {
	if existed, err := c.ctdb.Exist(c.db.WithContext(ctx), categoryID); err != nil || !existed {
		return ErrCategoryNotFound.SetInternal(err)
	}

	if err := c.ctdb.Delete(c.db.WithContext(ctx), categoryID); err != nil {
		return server.NewHTTPInternalError("Error deleting blog").SetInternal(err)
	}

	return nil
}
