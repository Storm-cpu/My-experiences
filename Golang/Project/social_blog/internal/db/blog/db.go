package blog

import (
	"social_blog/config"
	"social_blog/internal/model"
	dbutil "social_blog/pkg/util/db"

	"gorm.io/gorm"
)

// NewDB returns a new card database instance
func NewDB(cfg *config.Configuration) *DB {
	return &DB{dbutil.NewDB(&model.Blog{}), cfg}
}

type DB struct {
	*dbutil.DB
	cfg *config.Configuration
}

func (d *DB) FindBlogByUserID(db *gorm.DB, userID int) (*model.Blog, error) {
	rec := new(model.Blog)
	if err := d.View(db, rec, "user_id = ?", userID); err != nil {
		return nil, err
	}
	return rec, nil
}
