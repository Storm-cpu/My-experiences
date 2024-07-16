package user

import (
	"social_blog/config"
	"social_blog/internal/model"
	dbutil "social_blog/pkg/util/db"

	"gorm.io/gorm"
)

// NewDB returns a new card database instance
func NewDB(cfg *config.Configuration) *DB {
	return &DB{dbutil.NewDB(&model.User{}), cfg}
}

type DB struct {
	*dbutil.DB
	cfg *config.Configuration
}

func (d *DB) FindUserByUsername(db *gorm.DB, username string) (*model.User, error) {
	rec := new(model.User)
	if err := d.View(db, rec, "user_name = ?", username); err != nil {
		return nil, err
	}
	return rec, nil
}
