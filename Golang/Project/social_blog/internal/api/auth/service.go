package auth

import (
	"social_blog/internal/model"
	dbutil "social_blog/pkg/util/db"

	"gorm.io/gorm"
)

func New(db *gorm.DB, udb UserDB) *Auth {
	return &Auth{
		db:  db,
		udb: udb,
	}
}

type Auth struct {
	db  *gorm.DB
	udb UserDB
}

type UserDB interface {
	dbutil.Intf
	FindUserByUsername(db *gorm.DB, username string) (*model.User, error)
}
