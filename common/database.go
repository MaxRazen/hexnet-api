package common

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDbConnection(dbConfig DbConfig) *gorm.DB {
	var connection gorm.Dialector

	switch dbConfig.connection {
	case "sqlite":
		connection = sqlite.Open(dbConfig.path)
	default:
		panic("Fatal: DB Connection is not supported: " + dbConfig.connection)
	}

	db, err := gorm.Open(connection, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("Fatal: DB Connection could not be established: " + err.Error())
	}

	DB = db

	return DB
}

func InitTestDbConnection() *gorm.DB {
	connection := sqlite.Open(":memory:")

	db, err := gorm.Open(connection)

	if err != nil {
		panic("Fatal: DB Connection could not be established: " + err.Error())
	}

	DB = db

	return DB
}

func GetDB() *gorm.DB {
	if DB == nil {
		panic("The DB Connection is not initialized")
	}
	return DB
}
