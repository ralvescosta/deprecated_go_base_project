package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"markets/pkg/app/errors"
	"markets/pkg/app/interfaces"

	_ "github.com/lib/pq"
)

func Connect(logger interfaces.ILogger, shotdown chan bool) (*sql.DB, error) {
	connString, err := getConnectionString()
	if err != nil {
		logger.Error(fmt.Sprintf("[Database::Connect] - wrong database credentials %e", err))
		return nil, err
	}

	db, err := sql.Open("postgres", connString)
	if err != nil {
		logger.Error(fmt.Sprintf("[Database::Connect] - error while connect to database: %e", err))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Error(fmt.Sprintf("[Database::Connect] - error while check database connection: %v", err))
		return nil, err
	}

	secondsToSleep, err := strconv.Atoi(os.Getenv("DB_SECONDS_TO_PING"))
	if err != nil {
		logger.Error(fmt.Sprintf("[Database::Connect] - DB_SECONDS_TO_PING is required: %v", err))
		return nil, errors.NewInternalError(err.Error())
	}

	go signalShotdown(db, logger, secondsToSleep, shotdown)

	return db, nil
}

func getConnectionString() (string, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		return "", errors.NewInternalError("DB_HOST is required")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		return "", errors.NewInternalError("DB_PORT is required")
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		return "", errors.NewInternalError("DB_USER is required")
	}

	pass := os.Getenv("DB_PASSWORD")
	if pass == "" {
		return "", errors.NewInternalError("DB_PASSWORD is required")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		return "", errors.NewInternalError("DB_NAME is required")
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		pass,
		dbName,
	), nil
}

func signalShotdown(db *sql.DB, logger interfaces.ILogger, secondsToSleep int, shotdown chan bool) {
	time.Sleep(time.Duration(secondsToSleep) * time.Second)
	err := db.Ping()
	if err != nil {
		logger.Error(fmt.Sprintf("[Database::Connection] - Redis Connection failure : %s", err.Error()))
		shotdown <- true
	}
}
