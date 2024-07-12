package main

import (
	"social_blog/config"
	adminUser "social_blog/internal/api/v1/admin/user"
	userDB "social_blog/internal/db/user"
	dbutil "social_blog/pkg/util/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.LoadConfig()

	db := dbutil.ConnectDB(cfg)

	udb := userDB.NewDB(&cfg)

	adminUserSvc := adminUser.New(db, udb)

	e := echo.New()

	e.Use(middleware.Logger())

	v1aRouter := e.Group("/v1")
	v1aRouter = v1aRouter.Group("/admin")
	adminUser.NewHTTP(adminUserSvc, v1aRouter.Group("/users"))

	e.Logger.Fatal(e.Start(":8080"))
}
