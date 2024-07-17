package migration

import (
	"social_blog/config"
	"social_blog/internal/model"
	dbutil "social_blog/pkg/util/db"
)

func Run() {
	cfg := config.LoadConfig()

	db := dbutil.ConnectDB(cfg)

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Blog{})
	db.AutoMigrate(&model.Comment{})
}
