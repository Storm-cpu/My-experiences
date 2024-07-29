package main

import (
	"social_blog/config"

	adminAuth "social_blog/internal/api/v1/admin/auth"
	adminBlog "social_blog/internal/api/v1/admin/blog"
	adminCategory "social_blog/internal/api/v1/admin/category"
	adminComment "social_blog/internal/api/v1/admin/comment"
	adminUser "social_blog/internal/api/v1/admin/user"

	blogDB "social_blog/internal/db/blog"
	categoryDB "social_blog/internal/db/category"
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
	ctdb := categoryDB.NewDB(&cfg)

	crypterSvc := crypter.New()
	jwtAdminSvc := jwt.New(cfg.JwtAdminAlgorithm, cfg.JwtAdminSecret, cfg.JwtAdminDuration)

	authAdminSvc := adminAuth.New(db, udb, crypterSvc, jwtAdminSvc, &cfg)
	adminUserSvc := adminUser.New(db, udb, crypterSvc)
	adminBlogSvc := adminBlog.New(db, bdb)
	adminCommentSvc := adminComment.New(db, cdb)
	adminCategorySvc := adminCategory.New(db, ctdb)

	v1AdminRouter := e.Group("/v1")

	v1AuthenRouter := e.Group("/v1")
	adminAuth.NewHTTP(authAdminSvc, v1AuthenRouter.Group("/admin"))

	v1AdminRouter.Use(jwtAdminSvc.MWFunc())

	v1AdminRouter = v1AdminRouter.Group("/admin")
	adminUser.NewHTTP(adminUserSvc, v1AdminRouter.Group("/users"))
	adminBlog.NewHTTP(adminBlogSvc, v1AdminRouter.Group("/blogs"))
	adminComment.NewHTTP(adminCommentSvc, v1AdminRouter.Group("/comments"))
	adminCategory.NewHTTP(adminCategorySvc, v1AdminRouter.Group("/categories"))

	server.Start(e)
}
