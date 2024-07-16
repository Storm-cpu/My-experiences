package blog

import (
	"social_blog/internal/model"
	dbutil "social_blog/pkg/util/db"

	"gorm.io/gorm"
)

func New(db *gorm.DB, bdb MyDB) *Blog {
	return &Blog{
		db:  db,
		bdb: bdb,
	}
}

type Blog struct {
	db  *gorm.DB
	bdb MyDB
}

type MyDB interface {
	dbutil.Intf
	FindBlogByUserID(db *gorm.DB, userID int) (*model.Blog, error)
}
