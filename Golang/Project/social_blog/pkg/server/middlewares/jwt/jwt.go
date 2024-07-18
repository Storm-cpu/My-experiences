package jwt

import (
	"context"
	"fmt"
	"net/http"
	"social_blog/pkg/server"
	"social_blog/pkg/util/contextx"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

// New generates new JWT service necessery for auth middleware
func New(algo, secret string, duration int) *Service {
	signingMethod := jwt.GetSigningMethod(algo)
	if signingMethod == nil {
		panic("invalid jwt signing method")
	}
	return &Service{
		algo:     signingMethod,
		key:      []byte(secret),
		duration: time.Duration(duration) * time.Second,
	}
}

// Service provides a Json-Web-Token authentication implementation
type Service struct {
	key      []byte
	duration time.Duration
	algo     jwt.SigningMethod
}

func (j *Service) MWFunc(services ...*Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := j.ParseTokenFromHeader(c)
			if err != nil || !token.Valid {
				for _, svc := range services {
					if svc == nil {
						continue
					}
					t, svcErr := svc.ParseTokenFromHeader(c)
					if svcErr != nil || !t.Valid {
						continue
					}
					claims := t.Claims.(jwt.MapClaims)
					info := make(map[string]interface{})
					for key, val := range claims {
						c.Set(key, val)
						info[key] = val
					}
					ctx := c.Request().Context()
					ctx = context.WithValue(ctx, contextx.UserInfoKey, info)
					request := c.Request().WithContext(ctx)
					c.SetRequest(request)
					return next(c)
				}
				if err != nil {
					c.Logger().Errorf("error parsing token: %+v", err)
				}
				return server.NewHTTPError(http.StatusUnauthorized, "UNAUTHORIZED", "Your session is unauthorized or has expired.").SetInternal(err)
			}
			claims := token.Claims.(jwt.MapClaims)
			info := make(map[string]interface{})
			for key, val := range claims {
				c.Set(key, val)
				info[key] = val
			}
			ctx := c.Request().Context()
			ctx = context.WithValue(ctx, contextx.UserInfoKey, info)
			request := c.Request().WithContext(ctx)
			c.SetRequest(request)
			return next(c)
		}
	}
}

func (j *Service) ParseTokenFromHeader(c echo.Context) (*jwt.Token, error) {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return nil, fmt.Errorf("token not found")
	}
	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && strings.ToLower(parts[0]) == "bearer") {
		return nil, fmt.Errorf("token invalid")
	}

	return j.ParseToken(parts[1])
}

func (j *Service) ParseToken(input string) (*jwt.Token, error) {
	return jwt.Parse(input, func(token *jwt.Token) (interface{}, error) {
		if j.algo != token.Method {
			return nil, fmt.Errorf("token method mismatched")
		}
		return j.key, nil
	})
}

func (j *Service) GenerateToken(claims map[string]interface{}, expire *time.Time) (string, int, error) {
	if expire == nil {
		expTime := time.Now().Add(j.duration)
		expire = &expTime
	}
	claims["exp"] = expire.Unix()

	token := jwt.NewWithClaims(j.algo, jwt.MapClaims(claims))
	tokenString, err := token.SignedString(j.key)

	return tokenString, int(time.Until(*expire).Seconds()), err
}
