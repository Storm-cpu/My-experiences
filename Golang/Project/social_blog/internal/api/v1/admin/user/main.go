package user

import (
	"context"
	"net/http"
	"social_blog/internal/model"
	dbutil "social_blog/internal/util/db"

	"github.com/labstack/echo/v4"
)

func (u *User) Create(ctx context.Context, data CreatUserData) (*model.User, error) {
	rec := &model.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}

	if err := u.udb.Create(u.db.WithContext(ctx), rec); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return rec, nil
}

func (u *User) View(ctx context.Context, id int) (*model.User, error) {
	rec := new(model.User)
	if err := u.udb.View(u.db.WithContext(ctx), rec, id); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return rec, nil
}

func (u *User) List(ctx context.Context, lq *dbutil.ListQueryCondition, count *int64) ([]*model.User, error) {
	var data []*model.User
	if err := u.udb.List(u.db.WithContext(ctx), &data, lq, count); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return data, nil
}

func (u *User) Update(ctx context.Context, data UpdateUserData, userID int) (*model.User, error) {
	rec := &model.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}

	if err := u.udb.Update(u.db.WithContext(ctx), rec, userID); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := u.udb.View(u.db.WithContext(ctx), rec, userID); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return rec, nil
}

func (u *User) Delete(ctx context.Context, id int) error {

	if existed, err := u.udb.Exist(u.db.WithContext(ctx), id); err != nil || !existed {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := u.udb.Delete(u.db.WithContext(ctx), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil
}
