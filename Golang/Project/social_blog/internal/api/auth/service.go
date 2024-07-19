package auth

import (
	"social_blog/config"
	"social_blog/internal/model"
	dbutil "social_blog/pkg/util/db"
	"time"

	"gorm.io/gorm"
)

func New(db *gorm.DB, udb UserDB, cr Crypter, jwt JWT, cfg *config.Configuration) *Auth {
	return &Auth{
		db:  db,
		udb: udb,
		cr:  cr,
		jwt: jwt,
		cfg: cfg,
	}
}

type Auth struct {
	db  *gorm.DB
	udb UserDB
	cr  Crypter
	jwt JWT
	cfg *config.Configuration
}

type UserDB interface {
	dbutil.Intf
	FindUserByUsername(db *gorm.DB, username string) (*model.User, error)
	FindByRefreshToken(db *gorm.DB, token string) (*model.User, error)
}

type Crypter interface {
	CompareHashAndPassword(hasedPwd string, rawPwd string) bool
	GenRefreshToken(secret string) (string, error)
	ValidateRefreshToken(token, secret string) bool
	UID() string
}

type JWT interface {
	GenerateToken(map[string]interface{}, *time.Time) (string, int, error)
}
