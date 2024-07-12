package user

import (
	"social_blog/config"
	"social_blog/internal/model"
	dbutil "social_blog/pkg/util/db"
)

// NewDB returns a new card database instance
func NewDB(cfg *config.Configuration) *DB {
	return &DB{dbutil.NewDB(&model.User{}), cfg}
}

type DB struct {
	*dbutil.DB
	cfg *config.Configuration
}
