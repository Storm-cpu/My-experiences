package blog

import (
	"context"
	"net/http"
	"social_blog/internal/model"
	"social_blog/pkg/server"
	dbutil "social_blog/pkg/util/db"
	structutil "social_blog/pkg/util/struct"
)

var (
	ErrBlogNotFound = server.NewHTTPError(http.StatusBadRequest, "BLOG_NOTFOUND", "Blog không tồn tại")
)

// Create creates a new Blog
func (b *Blog) Create(ctx context.Context, data CreateBlogData) (*model.Blog, error) {
	rec := &model.Blog{
		Title:      data.Title,
		Content:    data.Content,
		AuthorID:   data.AuthorID,
		Status:     data.Status,
		Visibility: data.Visibility,
	}

	if err := b.bdb.Create(b.db.WithContext(ctx), rec); err != nil {
		return nil, server.NewHTTPInternalError("Error creating blog").SetInternal(err)
	}

	return rec, nil
}

// View returns a single Blog
func (b *Blog) View(ctx context.Context, id int) (*model.Blog, error) {
	rec := new(model.Blog)
	if err := b.bdb.View(b.db.WithContext(ctx), rec, id); err != nil {
		return nil, ErrBlogNotFound.SetInternal(err)
	}
	return rec, nil
}

// List returns a list of Blogs
func (b *Blog) List(ctx context.Context, lq *dbutil.ListQueryCondition, count *int64) ([]*model.Blog, error) {
	var data []*model.Blog
	if err := b.bdb.List(b.db.WithContext(ctx), &data, lq, count); err != nil {
		return nil, server.NewHTTPInternalError("Error listing blog").SetInternal(err)
	}
	return data, nil
}

// Update updates Blog information
func (b *Blog) Update(ctx context.Context, data UpdateBlogData, blogID int) (*model.Blog, error) {
	update := structutil.ToMap(data)
	if err := b.bdb.Update(b.db.WithContext(ctx), update, blogID); err != nil {
		return nil, server.NewHTTPInternalError("Error updating blog").SetInternal(err)
	}

	rec := new(model.Blog)
	if err := b.bdb.View(b.db.WithContext(ctx), rec, blogID); err != nil {
		return nil, ErrBlogNotFound.SetInternal(err)
	}

	return rec, nil
}

// Delete deletes a Blog
func (b *Blog) Delete(ctx context.Context, id int) error {
	if existed, err := b.bdb.Exist(b.db.WithContext(ctx), id); err != nil || !existed {
		return ErrBlogNotFound.SetInternal(err)
	}

	if err := b.bdb.Delete(b.db.WithContext(ctx), id); err != nil {
		return server.NewHTTPInternalError("Error deleting blog").SetInternal(err)
	}

	return nil
}
