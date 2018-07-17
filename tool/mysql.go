package tool

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	DBHost  = "127.0.0.1"
	DBPort  = ":3306"
	DBUser  = "root"
	DBPass  = ""
	DBDbase = "cms"
)

func InitializeDB() (*sql.DB, error) {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPass, DBHost, DBDbase)
	return sql.Open("mysql", dbConn)
}
