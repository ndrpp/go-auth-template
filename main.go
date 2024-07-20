package main

import (
	"go-auth-template/server"
	"go-auth-template/utils"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err.Error())
	}
	config := server.NewConfig("localhost", "8081")
	logger, err := utils.NewLogger()
	if err != nil {
		log.Fatal(err.Error())
	}

	dsn := os.Getenv("DATABASE_URL")
	db, err := utils.CreateDBClient(logger, dsn)
	if err != nil {
		logger.Error(err.Error())
	}

	s := server.NewServer(config, logger, db)
	err = server.Listen(s, logger)
	if err != nil {
		logger.Error(err.Error())
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
