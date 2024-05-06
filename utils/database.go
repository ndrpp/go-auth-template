package utils

import (
	"database/sql"
)

func CreateDBClient(logger *Logger, dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	if err = db.Ping(); err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return db, nil
}
