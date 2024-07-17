package auth

import (
	"social_blog/internal/model"
	dbutil "social_blog/pkg/util/db"

	"gorm.io/gorm"
)

func New(db *gorm.DB, udb UserDB, cr Crypter) *Auth {
	return &Auth{
		db:  db,
		udb: udb,
		cr:  cr,
	}
}

type Auth struct {
	db  *gorm.DB
	udb UserDB
	cr  Crypter
}

type UserDB interface {
	dbutil.Intf
	FindUserByUsername(db *gorm.DB, username string) (*model.User, error)
}

type Crypter interface {
	CompareHashAndPassword(hasedPwd string, rawPwd string) bool
}
