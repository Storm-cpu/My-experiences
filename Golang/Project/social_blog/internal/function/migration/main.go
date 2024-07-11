package migration

import (
	"social_blog/config"
	"social_blog/internal/model"
	dbutil "social_blog/internal/util/db"
)

func Run() {
	cfg := config.LoadConfig()

	db := dbutil.ConnectDB(cfg)

	db.AutoMigrate(&model.User{})
}
