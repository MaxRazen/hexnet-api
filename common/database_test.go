package common

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConnectingDatabase(t *testing.T) {
	asserts := assert.New(t)
	dbFile, _ := os.CreateTemp("", "app.db.*.sql")
	defer func() {
		_ = dbFile.Close()
		_ = os.Remove(dbFile.Name())
	}()

	dbConfig := DbConfig{
		connection: "sqlite",
		path:       dbFile.Name(),
	}
	db := InitDbConnection(dbConfig)
	// Test create & close DB
	_, err := os.Stat(dbConfig.path)
	asserts.NoError(err, "Db should exist")

	// Test get a connecting from connection pools
	if pinger, ok := db.ConnPool.(interface{ Ping() error }); ok {
		err := pinger.Ping()

		asserts.NoError(err, "Db should be able to ping")
	}
}

func TestInitTestDbConnection(t *testing.T) {
	asserts := assert.New(t)

	defer func() {
		if r := recover(); r != nil {
			asserts.NotNil(r, "Connection to test DB could not be established")
		}
	}()

	db := InitTestDbConnection()
	asserts.NotNil(db, "Test DB Connection could not be established")

	if pinger, ok := db.ConnPool.(interface{ Ping() error }); ok {
		err := pinger.Ping()

		asserts.NoError(err, "Db should be able to ping")
	}
}

func TestGetDB(t *testing.T) {
	asserts := assert.New(t)

	db := InitTestDbConnection()
	asserts.NotNil(db, "Test DB Connection could not be established")
	asserts.Same(db, GetDB(), "GetDB must return same instance")
}
