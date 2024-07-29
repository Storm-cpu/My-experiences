package auth

import (
	"context"
	"net/http"
	"social_blog/internal/model"
	"social_blog/pkg/server"
	"time"
)

var (
	ErrInvalidCredentials  = server.NewHTTPError(http.StatusUnauthorized, "INVALID_CREDENTIALS", "Username or password is incorrect")
	ErrInvalidRefreshToken = server.NewHTTPError(http.StatusUnauthorized, "INVALID_REFRESH_TOKEN", "Invalid refresh token")
	ErrGenerateToken       = server.NewHTTPInternalError("Error generating token")
)

func (s *Auth) Authenticate(ctx context.Context, data Credentials) (*model.AuthToken, error) {
	user, err := s.udb.FindUserByUsername(s.db.WithContext(ctx), data.Username)
	if err != nil || user == nil {
		return nil, ErrInvalidCredentials.SetInternal(err)
	}
	if !s.cr.CompareHashAndPassword(user.Password, data.Password) {
		return nil, ErrInvalidCredentials
	}
	if user.Blocked {
		return nil, server.NewHTTPError(http.StatusUnauthorized, "USER_BLOCKED", "Your account has been blocked and may not login")
	}

	return s.loginUser(user)
}

func (s *Auth) loginUser(u *model.User) (*model.AuthToken, error) {
	claims := map[string]interface{}{
		"id":       u.ID,
		"username": u.Username,
		"email":    u.Email,
	}
	token, expiresin, err := s.jwt.GenerateToken(claims, nil)
	if err != nil {
		return nil, ErrGenerateToken.SetInternal(err)
	}

	refreshToken, err := s.cr.GenRefreshToken(s.cfg.JwtAdminSecret)
	if err != nil {
		return nil, ErrGenerateToken.SetInternal(err)
	}

	err = s.udb.Update(s.db, map[string]interface{}{"refresh_token": refreshToken, "last_login": time.Now()}, u.ID)
	if err != nil {
		return nil, server.NewHTTPInternalError("Error updating user").SetInternal(err)
	}

	return &model.AuthToken{AccessToken: token, TokenType: "bearer", ExpiresIn: expiresin, RefreshToken: refreshToken}, nil
}

func (s *Auth) RefreshToken(ctx context.Context, data RefreshTokenData) (*model.AuthToken, error) {
	usr, err := s.udb.FindByRefreshToken(s.db.WithContext(ctx), data.RefreshToken)
	if err != nil || usr == nil {
		return nil, ErrInvalidRefreshToken.SetInternal(err)
	}
	if !s.cr.ValidateRefreshToken(usr.RefreshToken, s.cfg.JwtAdminSecret) {
		return nil, ErrInvalidRefreshToken
	}
	return s.loginUser(usr)
}
