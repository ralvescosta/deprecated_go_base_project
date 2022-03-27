package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"markets/pkg/app/errors"
	"markets/pkg/app/interfaces"

	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/pq"
)

var open = apmsql.Open

func Connect(logger interfaces.ILogger, shotdown chan bool) (*sql.DB, error) {
	connString, err := getConnectionString()
	if err != nil {
		logger.Error(fmt.Sprintf("[Database::Connect] - wrong database credentials %s", err.Error()))
		return nil, err
	}

	db, err := open("postgres", connString)
	if err != nil {
		logger.Error(fmt.Sprintf("[Database::Connect] - error while connect to database: %s", err.Error()))
		return nil, errors.NewInternalError(fmt.Sprintf("failure to connect to the database: %s", err.Error()))
	}

	err = db.Ping()
	if err != nil {
		logger.Error(fmt.Sprintf("[Database::Connect] - error while check database connection: %s", err.Error()))
		return nil, errors.NewInternalError(fmt.Sprintf("failure to connect to the database: %s", err.Error()))
	}

	secondsToSleep, err := strconv.Atoi(os.Getenv("DB_SECONDS_TO_PING"))
	if err != nil {
		logger.Error(fmt.Sprintf("[Database::Connect] - DB_SECONDS_TO_PING is required: %s", err.Error()))
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
		logger.Error(fmt.Sprintf("[Database::Connection] - Connection failure : %s", err.Error()))
		shotdown <- true
	}
}
