package comment

import (
	"context"
	"net/http"
	"social_blog/internal/model"
	"social_blog/pkg/server"
	dbutil "social_blog/pkg/util/db"
	structutil "social_blog/pkg/util/struct"
)

var (
	ErrCommentNotFound = server.NewHTTPError(http.StatusBadRequest, "COMMENT_NOTFOUND", "Comment không tồn tại")
)

// Create creates a new Comment
func (c *Comment) Create(ctx context.Context, data CreatCommentData) (*model.Comment, error) {
	rec := &model.Comment{
		BlogID:  data.BlogID,
		UserID:  data.UserID,
		Content: data.Content,
	}

	if err := c.cdb.Create(c.db.WithContext(ctx), rec); err != nil {
		return nil, server.NewHTTPInternalError("Error creating comment").SetInternal(err)
	}

	return rec, nil
}

// View returns a single Comment
func (c *Comment) View(ctx context.Context, id int) (*model.Comment, error) {
	rec := new(model.Comment)
	if err := c.cdb.View(c.db.WithContext(ctx), rec, id); err != nil {
		return nil, ErrCommentNotFound.SetInternal(err)
	}
	return rec, nil
}

// List returns a list of Users
func (c *Comment) List(ctx context.Context, lq *dbutil.ListQueryCondition, count *int64) ([]*model.Comment, error) {
	var data []*model.Comment
	if err := c.cdb.List(c.db.WithContext(ctx), &data, lq, count); err != nil {
		return nil, server.NewHTTPInternalError("Error listing comment").SetInternal(err)
	}
	return data, nil
}

// Update updates Comment information
func (c *Comment) Update(ctx context.Context, data UpdateCommentData, userID int) (*model.Comment, error) {
	update := structutil.ToMap(data)
	if err := c.cdb.Update(c.db.WithContext(ctx), update, userID); err != nil {
		return nil, server.NewHTTPInternalError("Error updating comment")
	}

	rec := new(model.Comment)
	if err := c.cdb.View(c.db.WithContext(ctx), rec, userID); err != nil {
		return nil, ErrCommentNotFound.SetInternal(err)
	}

	return rec, nil
}

// Delete deletes a Comment
func (c *Comment) Delete(ctx context.Context, id int) error {
	if existed, err := c.cdb.Exist(c.db.WithContext(ctx), id); err != nil || !existed {
		return ErrCommentNotFound.SetInternal(err)
	}

	if err := c.cdb.Delete(c.db.WithContext(ctx), id); err != nil {
		return server.NewHTTPInternalError("Error deleting comment").SetInternal(err)
	}

	return nil
}
