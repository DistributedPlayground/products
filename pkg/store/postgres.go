package store

import (
	"fmt"

	"github.com/DistributedPlayground/go-lib/common"
	"github.com/DistributedPlayground/products/config"
	"github.com/jmoiron/sqlx"
)

var pgDB *sqlx.DB
var DBDriver = "postgres"

func strConnection() string {
	var (
		DBUser     = config.Var.DB_USERNAME
		DBPassword = config.Var.DB_PASSWORD
		DBName     = config.Var.DB_NAME
		DBHost     = config.Var.DB_HOST
		DBPort     = config.Var.DB_PORT
	)

	var SSLMode string

	if common.IsLocalEnv() {
		SSLMode = "disable"
	} else {
		SSLMode = "require"
	}

	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		DBHost,
		DBPort,
		DBUser,
		DBPassword,
		DBName,
		SSLMode,
	)

	return str
}

func MustNewPG() *sqlx.DB {
	if pgDB != nil {
		return pgDB
	}

	pgDB, err := sqlx.Connect(DBDriver, strConnection())
	if err != nil {
		panic(err)
	}

	if err := pgDB.Ping(); err != nil {
		panic(err)
	}

	return pgDB
}
