package auth

import (
	"context"
	"net/http"
	"social_blog/pkg/server"
)

func (s *Auth) Authenticate(ctx context.Context, data Credentials) (*RepMessage, error) {
	user, err := s.udb.FindUserByUsername(s.db.WithContext(ctx), data.Username)
	if err != nil || user == nil {
		return nil, server.NewHTTPError(http.StatusUnauthorized, "INVALID_CREDENTIALS", "Username or password is incorrect")
	}
	if !s.cr.CompareHashAndPassword(user.Password, data.Password) {
		return nil, server.NewHTTPError(http.StatusUnauthorized, "INVALID_CREDENTIALS", "Username or password is incorrect")
	}
	if user.Blocked {
		return nil, server.NewHTTPError(http.StatusUnauthorized, "USER_BLOCKED", "Your account has been blocked and may not login")
	}

	return &RepMessage{Message: "Login success"}, nil
}
