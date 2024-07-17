package user

import (
	"context"
	"net/http"
	"social_blog/internal/model"
	"social_blog/pkg/server"
	dbutil "social_blog/pkg/util/db"
	structutil "social_blog/pkg/util/struct"
)

var (
	ErrUserNotFound = server.NewHTTPError(http.StatusBadRequest, "USER_NOTFOUND", "User không tồn tại")
)

// Create creates a new User
func (u *User) Create(ctx context.Context, data CreatUserData) (*model.User, error) {
	rec := &model.User{
		Username:    data.Username,
		Password:    u.cr.HashPassword(data.Password),
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		Blocked:     data.Blocked,
	}

	if err := u.udb.Create(u.db.WithContext(ctx), rec); err != nil {
		return nil, server.NewHTTPInternalError("Error creating user").SetInternal(err)
	}

	return rec, nil
}

// View returns a single User
func (u *User) View(ctx context.Context, id int) (*model.User, error) {
	rec := new(model.User)
	if err := u.udb.View(u.db.WithContext(ctx), rec, id); err != nil {
		return nil, ErrUserNotFound.SetInternal(err)
	}
	return rec, nil
}

// List returns a list of Users
func (u *User) List(ctx context.Context, lq *dbutil.ListQueryCondition, count *int64) ([]*model.User, error) {
	var data []*model.User
	if err := u.udb.List(u.db.WithContext(ctx), &data, lq, count); err != nil {
		return nil, server.NewHTTPInternalError("Error listing user").SetInternal(err)
	}
	return data, nil
}

// Update updates User information
func (u *User) Update(ctx context.Context, data UpdateUserData, userID int) (*model.User, error) {
	update := structutil.ToMap(data)
	if err := u.udb.Update(u.db.WithContext(ctx), update, userID); err != nil {
		return nil, server.NewHTTPInternalError("Error updating user")
	}

	rec := new(model.User)
	if err := u.udb.View(u.db.WithContext(ctx), rec, userID); err != nil {
		return nil, ErrUserNotFound.SetInternal(err)
	}

	return rec, nil
}

// Delete deletes a User
func (u *User) Delete(ctx context.Context, id int) error {
	if existed, err := u.udb.Exist(u.db.WithContext(ctx), id); err != nil || !existed {
		return ErrUserNotFound.SetInternal(err)
	}

	if err := u.udb.Delete(u.db.WithContext(ctx), id); err != nil {
		return server.NewHTTPInternalError("Error deleting user").SetInternal(err)
	}

	return nil
}
