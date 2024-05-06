package utils

import (
	"log"
	"os"

	"go.uber.org/zap"
)

type Logger struct {
	Logger *zap.Logger
}

func NewLogger() (*Logger, error) {
	env := os.Getenv("ENVIRONMENT")
	var logger *zap.Logger
	var err error
	if env == "production" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}

	return &Logger{Logger: logger}, nil
}

func (l *Logger) Info(message string) {
	l.Logger.Info(message)
}

func (l *Logger) Error(message string) {
	l.Logger.Error(message)
}

func (l *Logger) Warn(message string) {
	l.Logger.Warn(message)
}

func (l *Logger) Debug(message string) {
	l.Logger.Debug(message)
}
