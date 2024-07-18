package auth

import (
	"social_blog/internal/model"
	dbutil "social_blog/pkg/util/db"
	"time"

	"gorm.io/gorm"
)

func New(db *gorm.DB, udb UserDB, cr Crypter, jwt JWT) *Auth {
	return &Auth{
		db:  db,
		udb: udb,
		cr:  cr,
		jwt: jwt,
	}
}

type Auth struct {
	db  *gorm.DB
	udb UserDB
	cr  Crypter
	jwt JWT
}

type UserDB interface {
	dbutil.Intf
	FindUserByUsername(db *gorm.DB, username string) (*model.User, error)
}

type Crypter interface {
	CompareHashAndPassword(hasedPwd string, rawPwd string) bool
	UID() string
}

type JWT interface {
	GenerateToken(map[string]interface{}, *time.Time) (string, int, error)
}
