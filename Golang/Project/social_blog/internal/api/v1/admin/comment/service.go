package comment

import (
	dbutil "social_blog/pkg/util/db"

	"gorm.io/gorm"
)

func New(db *gorm.DB, cdb MyDB) *Comment {
	return &Comment{
		db:  db,
		cdb: cdb,
	}
}

type Comment struct {
	db  *gorm.DB
	cdb MyDB
}

type MyDB interface {
	dbutil.Intf
}
