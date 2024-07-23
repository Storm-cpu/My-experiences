package category

import (
	"social_blog/config"
	"social_blog/internal/model"
	dbutil "social_blog/pkg/util/db"
)

// NewDB returns a new comment database instance
func NewDB(cfg *config.Configuration) *DB {
	return &DB{dbutil.NewDB(&model.Category{}), cfg}
}

type DB struct {
	*dbutil.DB
	cfg *config.Configuration
}
