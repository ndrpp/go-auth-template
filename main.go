package main

import (
	"go-auth-template/server"
	"go-auth-template/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config := server.NewConfig("localhost", "8081")
	logger := utils.NewLogger()

	s := server.NewServer(config, logger)
	server.Listen(s, logger)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
