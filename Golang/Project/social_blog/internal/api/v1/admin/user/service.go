package user

import (
	dbutil "social_blog/pkg/util/db"

	"gorm.io/gorm"
)

func New(db *gorm.DB, udb MyDB, cr Crypter) *User {
	return &User{
		db:  db,
		udb: udb,
		cr:  cr,
	}
}

type User struct {
	db  *gorm.DB
	udb MyDB
	cr  Crypter
}

type MyDB interface {
	dbutil.Intf
}

type Crypter interface {
	HashPassword(string) string
}
