package main

import (
	"social_blog/config"

	"social_blog/internal/api/auth"
	adminBlog "social_blog/internal/api/v1/admin/blog"
	adminComment "social_blog/internal/api/v1/admin/comment"
	adminUser "social_blog/internal/api/v1/admin/user"

	blogDB "social_blog/internal/db/blog"
	commentDB "social_blog/internal/db/comment"
	userDB "social_blog/internal/db/user"

	"social_blog/pkg/server"
	"social_blog/pkg/server/middlewares/jwt"
	"social_blog/pkg/util/crypter"
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
	cdb := commentDB.NewDB(&cfg)

	crypterSvc := crypter.New()
	jwtAdminSvc := jwt.New(cfg.JwtAdminAlgorithm, cfg.JwtAdminSecret, cfg.JwtAdminDuration)

	authSvc := auth.New(db, udb, crypterSvc, jwtAdminSvc)
	adminUserSvc := adminUser.New(db, udb, crypterSvc)
	adminBlogSvc := adminBlog.New(db, bdb)
	adminCommentSvc := adminComment.New(db, cdb)

	auth.NewHTTP(authSvc, e)
	v1aRouter := e.Group("/v1")
	v1aRouter = v1aRouter.Group("/admin")
	adminUser.NewHTTP(adminUserSvc, v1aRouter.Group("/users"))
	adminBlog.NewHTTP(adminBlogSvc, v1aRouter.Group("/blogs"))
	adminComment.NewHTTP(adminCommentSvc, v1aRouter.Group("/comments"))

	server.Start(e)
}
