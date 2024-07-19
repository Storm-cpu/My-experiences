package main

import (
	"social_blog/config"

	adminAuth "social_blog/internal/api/v1/admin/auth"
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
	jwtSvc := jwt.New(cfg.JwtAlgorithm, cfg.JwtSecret, cfg.JwtDuration)

	authAdminSvc := adminAuth.New(db, udb, crypterSvc, jwtSvc, &cfg)
	adminUserSvc := adminUser.New(db, udb, crypterSvc)
	adminBlogSvc := adminBlog.New(db, bdb)
	adminCommentSvc := adminComment.New(db, cdb)

	v1aRouter := e.Group("/v1")
	v1aRouter = v1aRouter.Group("/admin")
	adminAuth.NewHTTP(authAdminSvc, v1aRouter)
	adminUser.NewHTTP(adminUserSvc, v1aRouter.Group("/users"))
	adminBlog.NewHTTP(adminBlogSvc, v1aRouter.Group("/blogs"))
	adminComment.NewHTTP(adminCommentSvc, v1aRouter.Group("/comments"))

	server.Start(e)
}
