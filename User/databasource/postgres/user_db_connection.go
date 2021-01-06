package userdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //using mySQL driver
)

const (
	mysqlUsername = "mysql_username"
	mysqlPassword = "mysql_password"
	mysqlHost     = "mysql_host"
	userSchema    = "mysql_schema"
)

var (
	//Client :- connection to the database
	Client   *sql.DB
	username = "root"
	password = "mysql"
	host     = "127.0.0.1:3306"
	schema   = "users"
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, schema)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Connected to the databse successfully")

}
