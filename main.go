package main

import (
	"database/sql"
	"go-auth-template/server"
	"go-auth-template/utils"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	config := server.NewConfig("localhost", "8081")
	logger := utils.NewLogger()
	err := godotenv.Load()
	if err != nil {
		logger.Error(err.Error())
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Error(err.Error())
	}

	s := server.NewServer(config, logger, db)
	server.Listen(s, logger)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
