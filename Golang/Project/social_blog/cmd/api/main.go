package main

import (
	"social_blog/config"
	adminUser "social_blog/internal/api/v1/admin/user"
	userDB "social_blog/internal/db/user"
	"social_blog/pkg/server"
	dbutil "social_blog/pkg/util/db"
)

func main() {
	cfg := config.LoadConfig()

	db := dbutil.ConnectDB(cfg)

	udb := userDB.NewDB(&cfg)

	e := server.New(&server.Config{
		Port:         cfg.Port,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})

	adminUserSvc := adminUser.New(db, udb)

	v1aRouter := e.Group("/v1")
	v1aRouter = v1aRouter.Group("/admin")
	adminUser.NewHTTP(adminUserSvc, v1aRouter.Group("/users"))

	e.Logger.Fatal(e.Start(":8080"))
}
