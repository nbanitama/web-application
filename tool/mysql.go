package tool

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbHost  = "127.0.0.1"
	dbPort  = ":3306"
	dbUser  = "root"
	dbPass  = ""
	dbDbase = "cms"
)

// to initialize DB connection
func InitializeDB() (*sql.DB, error) {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbDbase)
	return sql.Open("mysql", dbConn)
}
