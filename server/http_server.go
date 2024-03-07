package server

import (
	"context"
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

func NewServer(config *Config, logger *utils.Logger) *http.Server {
	mux := http.NewServeMux()
	ctx := context.Background()
	routes.AddBaseRoutes(ctx, mux, logger)
	var handler http.Handler = mux
	handler = middleware.CorsMiddeware(handler)

	return &http.Server{
		Addr:         fmt.Sprintf("%s:%s", config.Host, config.Port),
		Handler:      handler,
		WriteTimeout: time.Hour,
		ReadTimeout:  time.Hour,
	}
}

func Listen(s *http.Server, logger *utils.Logger) error {
	logger.Info(fmt.Sprintf("Server is listening on: %s", s.Addr))
	return s.ListenAndServe()
}
