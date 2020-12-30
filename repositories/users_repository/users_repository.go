package users_repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/black-lotus/bookstore-utils-go/logger"
	"github.com/go-sql-driver/mysql"
)

const (
	mysqlUsername = "mysql_username"
	mysqlPassword = "mysql_password"
	mysqlHost     = "mysql_host"
	mysqlDatabase = "mysql_database"
)

var (
	Client *sql.DB

	username = os.Getenv(mysqlUsername)
	password = os.Getenv(mysqlPassword)
	host     = os.Getenv(mysqlHost)
	database = os.Getenv(mysqlDatabase)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, database,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured")
}
