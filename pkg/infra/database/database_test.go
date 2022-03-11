package database

import (
	"database/sql"
	"errors"
	"os"
	"testing"

	"markets/pkg/infra/logger"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func Test_Connect(t *testing.T) {
	t.Run("should connect to database correctly", func(t *testing.T) {
		sut := makeDatabaseSutRtn(nil)

		db, err := Connect(sut.logger, sut.shotdown)

		assert.NotNil(t, db)
		assert.NoError(t, err)
	})

	t.Run("should return error if some error when try to connect", func(t *testing.T) {
		sut := makeDatabaseSutRtn(errors.New("some error"))
		sut.logger.On("Error", "[Database::Connect] - error while connect to database: some error", []zapcore.Field(nil))

		db, err := Connect(sut.logger, sut.shotdown)

		assert.Nil(t, db)
		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should return error if DB_SECONDS_TO_PING has not been defined", func(t *testing.T) {
		sut := makeDatabaseSutRtn(nil)
		os.Setenv("DB_SECONDS_TO_PING", "")
		sut.logger.On("Error", "[Database::Connect] - DB_SECONDS_TO_PING is required: strconv.Atoi: parsing \"\": invalid syntax", []zapcore.Field(nil))

		db, err := Connect(sut.logger, sut.shotdown)

		assert.Nil(t, db)
		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})
}

func Test_GetConnectionString(t *testing.T) {
	t.Run("should return err if DB_HOST was has not been defined", func(t *testing.T) {
		sut := makeDatabaseSutRtn(nil)
		os.Setenv("DB_HOST", "")
		sut.logger.On("Error", "[Database::Connect] - wrong database credentials DB_HOST is required", []zapcore.Field(nil))

		db, err := Connect(sut.logger, sut.shotdown)

		assert.Nil(t, db)
		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should return err if DB_PORT was has not been defined", func(t *testing.T) {
		sut := makeDatabaseSutRtn(nil)
		os.Setenv("DB_PORT", "")
		sut.logger.On("Error", "[Database::Connect] - wrong database credentials DB_PORT is required", []zapcore.Field(nil))

		db, err := Connect(sut.logger, sut.shotdown)

		assert.Nil(t, db)
		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should return err if DB_USER was has not been defined", func(t *testing.T) {
		sut := makeDatabaseSutRtn(nil)
		os.Setenv("DB_USER", "")
		sut.logger.On("Error", "[Database::Connect] - wrong database credentials DB_USER is required", []zapcore.Field(nil))

		db, err := Connect(sut.logger, sut.shotdown)

		assert.Nil(t, db)
		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should return err if DB_PASSWORD was has not been defined", func(t *testing.T) {
		sut := makeDatabaseSutRtn(nil)
		os.Setenv("DB_PASSWORD", "")
		sut.logger.On("Error", "[Database::Connect] - wrong database credentials DB_PASSWORD is required", []zapcore.Field(nil))

		db, err := Connect(sut.logger, sut.shotdown)

		assert.Nil(t, db)
		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})

	t.Run("should return err if DB_NAME was has not been defined", func(t *testing.T) {
		sut := makeDatabaseSutRtn(nil)
		os.Setenv("DB_NAME", "")
		sut.logger.On("Error", "[Database::Connect] - wrong database credentials DB_NAME is required", []zapcore.Field(nil))

		db, err := Connect(sut.logger, sut.shotdown)

		assert.Nil(t, db)
		assert.Error(t, err)
		sut.logger.AssertExpectations(t)
	})
}

type databaseSutRtn struct {
	logger   *logger.LoggerSpy
	shotdown chan bool
}

func makeDatabaseSutRtn(dbConnError error) databaseSutRtn {
	os.Setenv("DB_HOST", "host")
	os.Setenv("DB_PORT", "1111")
	os.Setenv("DB_USER", "user")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "name")
	os.Setenv("DB_SECONDS_TO_PING", "20")

	open = func(driver, connectionString string) (*sql.DB, error) {
		db, _, _ := sqlmock.New()
		return db, dbConnError
	}

	logger := logger.NewLoggerSpy()
	shotdown := make(chan bool)

	return databaseSutRtn{logger, shotdown}
}
