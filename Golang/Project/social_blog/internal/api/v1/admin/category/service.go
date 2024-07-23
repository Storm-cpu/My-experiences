package category

import (
	dbutil "social_blog/pkg/util/db"

	"gorm.io/gorm"
)

func New(db *gorm.DB, ctdb MyDB) *Category {
	return &Category{
		db:   db,
		ctdb: ctdb,
	}
}

type Category struct {
	db   *gorm.DB
	ctdb MyDB
}

type MyDB interface {
	dbutil.Intf
}
