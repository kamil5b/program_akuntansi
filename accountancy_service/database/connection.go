package database

import (
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectManual(database_type, url, user, password, protocol, database string, models ...interface{}) {

	dsn := user + ":" + password + "@" + protocol + "(" + url + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	ConnectDSN(database_type, dsn, models...)
}

func ConnectDSN(database_type, dsn string, models ...interface{}) {
	var dialect gorm.Dialector

	database_type = strings.ToLower(database_type)

	switch database_type {
	case "mysql":
		dialect = mysql.Open(dsn)
	case "postgres":
		dialect = postgres.Open(dsn)
	case "sqlserver":
		dialect = sqlserver.Open(dsn)
	case "sqlite":
		dialect = sqlite.Open(dsn)
	default:
		panic("not valid database type")
	}

	connection, err := gorm.Open(dialect, &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection
	connection.AutoMigrate(models...)
}
