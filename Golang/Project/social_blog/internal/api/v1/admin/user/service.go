package user

import (
	dbutil "social_blog/pkg/util/db"

	"gorm.io/gorm"
)

func New(db *gorm.DB, udb MyDB) *User {
	return &User{
		db:  db,
		udb: udb,
	}
}

type User struct {
	db  *gorm.DB
	udb MyDB
}

type MyDB interface {
	dbutil.Intf
}
