package main

import (
	"social_blog/config"
	adminBlog "social_blog/internal/api/v1/admin/blog"
	adminUser "social_blog/internal/api/v1/admin/user"
	blogDB "social_blog/internal/db/blog"
	userDB "social_blog/internal/db/user"
	"social_blog/pkg/server"
	dbutil "social_blog/pkg/util/db"
)

func main() {
	cfg := config.LoadConfig()

	db := dbutil.ConnectDB(cfg)

	e := server.New(&server.Config{
		Port:         cfg.Port,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})

	udb := userDB.NewDB(&cfg)
	bdb := blogDB.NewDB(&cfg)

	adminUserSvc := adminUser.New(db, udb)
	adminBlogSvc := adminBlog.New(db, bdb)

	v1aRouter := e.Group("/v1")
	v1aRouter = v1aRouter.Group("/admin")
	adminUser.NewHTTP(adminUserSvc, v1aRouter.Group("/users"))
	adminBlog.NewHTTP(adminBlogSvc, v1aRouter.Group("/blogs"))

	server.Start(e)
}
