package server

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	Port         int
	ReadTimeout  int
	WriteTimeout int
}

// DefaultConfig for the API server
var DefaultConfig = Config{
	Port:         8080,
	ReadTimeout:  10,
	WriteTimeout: 5,
}

func (c *Config) fillDefaults() {
	if c.Port == 0 {
		c.Port = DefaultConfig.Port
	}
	if c.ReadTimeout == 0 {
		c.ReadTimeout = DefaultConfig.ReadTimeout
	}
	if c.WriteTimeout == 0 {
		c.WriteTimeout = DefaultConfig.WriteTimeout
	}
}

// New instantates new Echo server
func New(cfg *Config) *echo.Echo {
	cfg.fillDefaults()
	e := echo.New()

	e.Use(middleware.Logger())

	e.Server.Addr = fmt.Sprintf(":%d", cfg.Port)
	e.Validator = NewValidator()
	e.HTTPErrorHandler = NewErrorHandler(e).Handle
	e.Server.ReadTimeout = time.Duration(cfg.ReadTimeout) * time.Minute
	e.Server.WriteTimeout = time.Duration(cfg.WriteTimeout) * time.Minute

	return e
}
