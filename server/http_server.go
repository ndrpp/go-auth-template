package server

import (
	"database/sql"
	"fmt"
	"go-auth-template/middleware"
	"go-auth-template/routes"
	"go-auth-template/utils"
	"net/http"
	"time"
)

type Config struct {
	Host string
	Port string
}

func NewConfig(host, port string) *Config {
	return &Config{
		Host: host,
		Port: port,
	}
}

func NewServer(config *Config, logger *utils.Logger, db *sql.DB) *http.Server {
	mux := http.NewServeMux()
	routes.AddBaseRoutes(mux, logger)
	routes.AddAuthRoutes(mux, logger, db)
	var handler http.Handler = mux
	handler = middleware.CorsMiddeware(handler)

	return &http.Server{
		Addr:              fmt.Sprintf("%s:%s", config.Host, config.Port),
		Handler:           http.TimeoutHandler(handler, time.Second, "Timeout"),
		ReadTimeout:       500 * time.Millisecond,
		ReadHeaderTimeout: 500 * time.Millisecond,
		IdleTimeout:       time.Second,
	}
}

func Listen(s *http.Server, logger *utils.Logger) error {
	logger.Info(fmt.Sprintf("Server is listening on: %s", s.Addr))
	return s.ListenAndServe()
}
