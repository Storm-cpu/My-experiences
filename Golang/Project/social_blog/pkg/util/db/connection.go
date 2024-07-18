package dbutil

import (
	"fmt"
	"log"
	"social_blog/config"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db       *gorm.DB
	instance bool
	lock     sync.Mutex
)

func ConnectDB(cfg config.Configuration) *gorm.DB {
	lock.Lock()
	defer lock.Unlock()

	if instance {
		fmt.Println("Database already connected!")
		return db
	}

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.PostgresHost, cfg.PostgresUser, cfg.PostgresDB, cfg.PostgresSSLMode, cfg.PostgresPassword)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("Connected to Database!")
	instance = true

	return db
}
